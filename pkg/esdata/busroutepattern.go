package esdata

import (
	"time"

	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

type BusRoutePattern struct {
	Date      time.Time `json:"date"`
	Title     string    `json:"title"`
	Kana      string    `json:"kana"`
	Operator  string    `json:"operator"`
	Route     string    `json:"route"`
	Pattern   string    `json:"pattern"`
	Location  *Location `json:"location,omitempty"`
	Next      *Location `json:"nextLocation,omitempty"`
	Direction string    `json:"direction"`
}

func FromODPTBusRoutePattern(bsr []*odpt.BusRoutePattern, poles odpt.BusStopPoleLookup) []*BusRoutePattern {
	esbsr := make([]*BusRoutePattern, len(bsr))[:0]

	for _, entry := range bsr {
		for i, poleOrder := range entry.PoleOrder {
			pole := poles[poleOrder.Pole]
			untypedRoute := removeType(entry.Route)
			route := removeFirstPeriodSeparatedChunk(untypedRoute)
			pattern := &BusRoutePattern{
				Date:     entry.Date.Date,
				Title:    entry.Title,
				Kana:     entry.Kana,
				Operator: removeType(entry.Operator),
				Route:    removeType(route),
				Pattern:  entry.Pattern,
			}

			if pole.Latitude != 0 {
				pattern.Location = &Location{
					Latitude:  pole.Latitude,
					Longitude: pole.Longitude,
				}
			}

			if i < len(entry.PoleOrder)-1 {
				nextPole := poles[entry.PoleOrder[i+1].Pole]
				pattern.Next = &Location{
					Latitude:  nextPole.Latitude,
					Longitude: nextPole.Longitude,
				}
			}

			esbsr = append(esbsr, pattern)
		}
	}

	return esbsr
}
