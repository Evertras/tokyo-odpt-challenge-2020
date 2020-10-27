package odpt

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Station struct {
	Base
	Location
	Date       time.Time `json:"dc:date"`
	Title      string    `json:"dc:title"`
}

type StationLookup map[string]*Station

// LoadStationsJSON loads all Station entries from a static JSON
// file created by the data dump API
func LoadStationsJSON(filename string) ([]*Station, error) {
	f, err := os.Open("./data/Station.json")

	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	stations := []*Station{}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&stations)

	if err != nil {
		return nil, fmt.Errorf("decoder.Decode: %w", err)
	}

	return stations, nil
}

// NewStationLookup creates a new lookup table to find station data
func NewStationLookup(stations []*Station) StationLookup {
	lookup := make(map[string]*Station, len(stations))

	for _, station := range stations {
		lookup[station.SameAs] = station
	}

	return lookup
}
