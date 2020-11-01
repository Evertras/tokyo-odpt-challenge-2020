package esdata

import (
	"time"

	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

type BusRoutePattern struct {
	Date     time.Time `json:"date"`
	Title    string    `json:"title"`
	Kana     string    `json:"kana"`
	Operator string    `json:"operator"`
	Route    string    `json:"route"`
	Pattern  string    `json:"pattern"`
	Location Location  `json:"location"`
}

func FromODPTBusRoutePattern(bsr []*odpt.BusRoutePattern, poles odpt.BusStopPoleLookup) []*BusRoutePattern {
	esbsr := make([]*BusRoutePattern, len(bsr))[:0]

	for _, entry := range bsr {
		for _, poleOrder := range entry.PoleOrder {
			pole := poles[poleOrder.Pole]
			untypedRoute := removeType(entry.Route)
			route := removeFirstPeriodSeparatedChunk(untypedRoute)
			esbsr = append(esbsr, &BusRoutePattern{
				Date:     entry.Date.Date,
				Title:    entry.Title,
				Kana:     entry.Kana,
				Operator: removeType(entry.Operator),
				Route:    removeType(route),
				Pattern:  entry.Pattern,
				Location: Location{
					Latitude:  pole.Latitude,
					Longitude: pole.Longitude,
				},
			})
		}
	}

	return esbsr
}
