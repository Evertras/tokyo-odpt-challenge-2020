package odpt

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dhconnelly/rtreego"
)

type BusStopPoleLookup map[string]*BusStopPole

// BusStopPole contains information about a bus stop in the world
// that may serve one or more buses
type BusStopPole struct {
	Base
	Location
	Valid

	Title          string            `json:"dc:title"`
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

func NewBusStopPoleSpatialTree(poles []*BusStopPole) *rtreego.Rtree {
	tree := rtreego.NewTree(2, 25, 50)

	for _, pole := range poles {
		if pole.Latitude > 0 {
			tree.Insert(pole)
		}
	}

	return tree
}
