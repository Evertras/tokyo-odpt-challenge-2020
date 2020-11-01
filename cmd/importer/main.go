package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v7"

	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/esdata"
	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

func main() {
	ctx := context.TODO()

	///////////////////////////////////////////////////////////////////////////
	// Load the data
	ps, err := odpt.LoadPassengerSurveysJSON("./data/PassengerSurvey.json")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Got %d passenger surveys", len(ps)))

	stations, err := odpt.LoadStationsJSON("./data/Stations.json")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Got %d stations", len(stations)))
	log.Println(fmt.Sprintf("%+v", stations[0]))

	stationLookup := odpt.NewStationLookup(stations)
	log.Println(fmt.Sprintf("%+v", stationLookup["odpt.Station:JR-East.ChuoRapid.Shinjuku"]))

	bsp, err := odpt.LoadBusStopPoleJSON("./data/BusstopPole.json")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Got %d bus stop poles", len(bsp))
	log.Println(fmt.Sprintf("%+v", bsp[0]))

	busStopPoleLookup := odpt.NewBusStopPoleLookup(bsp)

	bsr, err := odpt.LoadBusRoutePatternJSON("./data/BusroutePattern.json")

	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("TOKEN")

	if token == "" {
		log.Fatal("Need to set TOKEN")
	}

	client := odpt.NewClient(token)

	b, err := client.GetAllBuses(ctx)

	if err != nil {
		log.Fatal("Failed to get all buses:", err)
	}

	log.Printf("%d buses currently running", len(b))

	///////////////////////////////////////////////////////////////////////////
	// Prepare ES
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

	///////////////////////////////////////////////////////////////////////////
	// Do the importing
	importer := esdata.NewImporter(es, stationLookup, busStopPoleLookup)

	err = importer.DeleteAllDataIndices()
	if err != nil {
		log.Fatal(err)
	}

	err = importer.ImportPassengerSurvey(ctx, ps)
	if err != nil {
		log.Fatal(err)
	}

	err = importer.ImportBusStopPole(ctx, bsp)
	if err != nil {
		log.Fatal(err)
	}

	err = importer.ImportBusRoutePattern(ctx, bsr)
	if err != nil {
		log.Fatal(err)
	}

	err = importer.ImportBus(ctx, b)
}
