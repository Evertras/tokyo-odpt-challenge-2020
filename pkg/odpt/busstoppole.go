package odpt

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// BusstopPole contains information about a bus stop in the world
// that may serve one or more buses
type BusstopPole struct {
	Base
	Date
	Location

	Title string `json:"dc:title"`
	Valid time.Time `json:"dct:valid"`
	Kana string `json:"odpt:kana"`
	TitleLocalized map[string]string `json:"title"`
	Operator []string `json:"odpt:operator"`
}

// LoadBusstopPoleJSON loads all BusstopPole entries from a static JSON
// file created by the data dump API
func LoadBusstopPoleJSON(filename string) ([]*BusstopPole, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	bsp := []*BusstopPole{}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&bsp)

	if err != nil {
		return nil, fmt.Errorf("decoder.Decode: %w", err)
	}

	return bsp, nil
}
