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

func ImportPassengerSurvey(ctx context.Context, esClient *elasticsearch.Client, ps []*odpt.PassengerSurvey) error {
	converted := FromODPTPassengerSurvey(ps)

	bulk, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index: IndexNamePassengerSurvey,
		Client: esClient,
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
			Body:  bytes.NewReader(data),
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
