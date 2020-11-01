package odpt

import (
	"encoding/json"
	"fmt"
	"os"
)

// BusRoutePattern describes a specific bus route
type BusRoutePattern struct {
	Base
	Valid

	Title     string             `json:"dc:title"`
	Kana      string             `json:"odpt:kana"`
	Operator  string             `json:"odpt:operator"`
	Route     string             `json:"odpt:busroute"`
	Pattern   string             `json:"odpt:pattern"`
	Direction string             `json:"odpt:direction"`
	Note      string             `json:"odpt:note"`
	PoleOrder []BusStopPoleOrder `json:"odpt:busstopPoleOrder"`
}

// LoadBusRoutePatternJSON loads all BusRoutePattern entries from a static JSON
// file created by the data dump API
func LoadBusRoutePatternJSON(filename string) ([]*BusRoutePattern, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	brp := []*BusRoutePattern{}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&brp)

	if err != nil {
		return nil, fmt.Errorf("decoder.Decode: %w", err)
	}

	return brp, nil
}
