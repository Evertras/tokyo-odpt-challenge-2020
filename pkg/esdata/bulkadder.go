package esdata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

type bulkAdder struct {
	indexer esutil.BulkIndexer
}

func startBulkAdder(ctx context.Context, esClient *elasticsearch.Client, index string) (*bulkAdder, error) {
	bulk, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  index,
		Client: esClient,
		OnError: func(ctx context.Context, err error) {
			log.Println("ERR:", err)
		},
	})

	if err != nil {
		return nil, fmt.Errorf("esutil.NewBulkIndexer: %w", err)
	}

	return &bulkAdder{
		indexer: bulk,
	}, nil
}

func (b *bulkAdder) add(ctx context.Context, data interface{}) error {
	marshalled, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = b.indexer.Add(ctx, esutil.BulkIndexerItem{
		Action: "index",
		Body:   bytes.NewReader(marshalled),
	})

	if err != nil {
		return fmt.Errorf("b.indexer.Add: %w", err)
	}

	return nil
}

// closeWithLoggedError eats any error and just logs it out for
// easy defers where we want to drop the error anyway, but still
// want to see in logs if something weird happened
func (b *bulkAdder) closeWithLoggedError(ctx context.Context) {
	if b == nil || b.indexer == nil {
		log.Println("Bulk indexer nil")
		return
	}

	err := b.indexer.Close(ctx)

	if err != nil {
		log.Println("Bulk indexer failed to close:", err)
	}
}
