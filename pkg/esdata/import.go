package esdata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

type Importer struct {
	esClient      *elasticsearch.Client
	stationLookup odpt.StationLookup
}

func NewImporter(esClient *elasticsearch.Client, stationLookup odpt.StationLookup) *Importer {
	return &Importer{
		esClient:      esClient,
		stationLookup: stationLookup,
	}
}

func (i *Importer) ImportPassengerSurvey(ctx context.Context, ps []*odpt.PassengerSurvey) error {
	converted := FromODPTPassengerSurvey(ps, i.stationLookup)

	_, err := i.esClient.Indices.Delete([]string{IndexNamePassengerSurvey})

	if err != nil {
		return fmt.Errorf("esapi.IndicesDelete: %w", err)
	}

	_, err = i.esClient.Indices.Create(IndexNamePassengerSurvey)

	if err != nil {
		return fmt.Errorf("esapi.IndicesCreate: %w", err)
	}

	indexMappingBody, err := genIndexMappingBody(map[string]string{
		"location": "geo_point",
	})

	res, err := i.esClient.Indices.PutMapping(
		indexMappingBody,
		i.esClient.Indices.PutMapping.WithIndex(IndexNamePassengerSurvey),
	)

	if err != nil {
		return fmt.Errorf("creating index with mapping: %w", err)
	}

	if res.StatusCode != 200 {
		panic(res)
	}

	bulk, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  IndexNamePassengerSurvey,
		Client: i.esClient,
		OnError: func(ctx context.Context, err error) {
			log.Println("ERR:", err)
		},
	})

	if err != nil {
		return fmt.Errorf("esutil.NewBulkIndexer: %w", err)
	}

	for i, entry := range converted {
		data, err := json.Marshal(entry)

		if err != nil {
			return fmt.Errorf("json.Marshal #%d: %w", i, err)
		}

		err = bulk.Add(ctx, esutil.BulkIndexerItem{
			Action: "index",
			Body:   bytes.NewReader(data),
		})

		if err != nil {
			bulk.Close(ctx)
			return fmt.Errorf("bulk.Close: %w", err)
		}
	}

	err = bulk.Close(ctx)

	if err != nil {
		return fmt.Errorf("bulk.Close: %w", err)
	}

	return nil
}
