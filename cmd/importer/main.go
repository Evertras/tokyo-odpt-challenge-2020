package main

import (
	"context"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"

	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/esdata"
	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

func main() {
	ctx := context.TODO()

	ps, err := odpt.LoadPassengerSurveysJSON("./data/PassengerSurvey.json")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Got %d: sample %+v", len(ps), ps[0]))

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)
	res.Body.Close()

	err = esdata.ImportPassengerSurvey(ctx, es, ps)

	if err != nil {
		log.Fatal(err)
	}
}
