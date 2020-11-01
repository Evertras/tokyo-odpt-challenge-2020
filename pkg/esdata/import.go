package esdata

import (
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

type Importer struct {
	esClient          *elasticsearch.Client
	stationLookup     odpt.StationLookup
	busStopPoleLookup odpt.BusStopPoleLookup
}

func NewImporter(
	esClient *elasticsearch.Client,
	stationLookup odpt.StationLookup,
	busStopPoleLookup odpt.BusStopPoleLookup) *Importer {
	return &Importer{
		esClient:          esClient,
		stationLookup:     stationLookup,
		busStopPoleLookup: busStopPoleLookup,
	}
}

func (i *Importer) DeleteAllDataIndices() error {
	toDelete := []string{
		IndexNameBusRoutePattern,
		IndexNameBusStopPole,
		IndexNamePassengerSurvey,
	}

	for _, index := range toDelete {
		// Do these separately since they may not all exist
		res, err := i.esClient.Indices.Delete([]string{index})

		if err != nil {
			return fmt.Errorf("i.esClient.Indices.Delete %q: %w", index, err)
		}

		if res.StatusCode/200 != 1 && res.StatusCode != 404 {
			return fmt.Errorf("unexpected status code for %q: %d", index, res.StatusCode)
		}
	}

	return nil
}

func (i *Importer) ImportPassengerSurvey(ctx context.Context, ps []*odpt.PassengerSurvey) error {
	err := i.prepLocationMapping(IndexNamePassengerSurvey, []string{"location"})

	if err != nil {
		return fmt.Errorf("i.prepLocationMapping: %w", err)
	}

	bulk, err := startBulkAdder(ctx, i.esClient, IndexNamePassengerSurvey)
	if err != nil {
		return fmt.Errorf("startBulkAdder: %w", err)
	}
	defer bulk.closeWithLoggedError(ctx)

	converted := FromODPTPassengerSurvey(ps, i.stationLookup)
	for i, entry := range converted {
		err = bulk.add(ctx, entry)

		if err != nil {
			return fmt.Errorf("bulk.add #%d: %w", i, err)
		}
	}

	return nil
}

func (i *Importer) ImportBusStopPole(ctx context.Context, bsp []*odpt.BusStopPole) error {
	err := i.prepLocationMapping(IndexNameBusStopPole, []string{"location"})
	if err != nil {
		return fmt.Errorf("i.prepLocationMapping: %w", err)
	}

	bulk, err := startBulkAdder(ctx, i.esClient, IndexNameBusStopPole)
	if err != nil {
		return fmt.Errorf("startBulkAdder: %w", err)
	}
	defer bulk.closeWithLoggedError(ctx)

	converted := FromODPTBusStopPole(bsp)
	for i, entry := range converted {
		err = bulk.add(ctx, entry)

		if err != nil {
			return fmt.Errorf("bulk.add #%d: %w", i, err)
		}
	}

	return nil
}

func (i *Importer) ImportBusRoutePattern(ctx context.Context, bsr []*odpt.BusRoutePattern) error {
	err := i.prepLocationMapping(IndexNameBusRoutePattern, []string{"location", "nextLocation"})
	if err != nil {
		return fmt.Errorf("i.prepLocationMapping: %w", err)
	}

	bulk, err := startBulkAdder(ctx, i.esClient, IndexNameBusRoutePattern)
	if err != nil {
		return fmt.Errorf("startBulkAdder: %w", err)
	}
	defer bulk.closeWithLoggedError(ctx)

	converted := FromODPTBusRoutePattern(bsr, i.busStopPoleLookup)
	for i, entry := range converted {
		err = bulk.add(ctx, entry)

		if err != nil {
			return fmt.Errorf("bulk.add #%d: %w", i, err)
		}
	}

	return nil
}

func (i *Importer) prepLocationMapping(index string, locFields []string) error {
	_, err := i.esClient.Indices.Delete([]string{index})

	if err != nil {
		return fmt.Errorf("esapi.IndicesDelete: %w", err)
	}

	_, err = i.esClient.Indices.Create(index)

	if err != nil {
		return fmt.Errorf("esapi.IndicesCreate: %w", err)
	}

	mapping := make(map[string]string)
	for _, field := range locFields {
		mapping[field] = "geo_point"
	}
	indexMappingBody, err := genIndexMappingBody(mapping)

	res, err := i.esClient.Indices.PutMapping(
		indexMappingBody,
		i.esClient.Indices.PutMapping.WithIndex(index),
	)

	if err != nil {
		return fmt.Errorf("creating index with mapping: %w", err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("unexpected error code %d", res.StatusCode)
	}

	return nil
}
