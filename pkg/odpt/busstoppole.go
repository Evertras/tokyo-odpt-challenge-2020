package odpt

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type BusStopPoleLookup map[string]*BusStopPole

// BusStopPole contains information about a bus stop in the world
// that may serve one or more buses
type BusStopPole struct {
	Base
	Location

	Title          string            `json:"dc:title"`
	Valid          time.Time         `json:"dct:valid"`
	Kana           string            `json:"odpt:kana"`
	TitleLocalized map[string]string `json:"title"`
	Operator       []string          `json:"odpt:operator"`
}

// LoadBusStopPoleJSON loads all BusStopPole entries from a static JSON
// file created by the data dump API
func LoadBusStopPoleJSON(filename string) ([]*BusStopPole, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	bsp := []*BusStopPole{}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&bsp)

	if err != nil {
		return nil, fmt.Errorf("decoder.Decode: %w", err)
	}

	return bsp, nil
}

// NewBusStopPoleLookup creates a new lookup table to find station data
func NewBusStopPoleLookup(poles []*BusStopPole) BusStopPoleLookup {
	lookup := make(map[string]*BusStopPole, len(poles))

	for _, pole := range poles {
		lookup[pole.SameAs] = pole
	}

	return lookup
}
