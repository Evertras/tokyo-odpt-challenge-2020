package esdata

import (
	"time"

	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

type BusStopPole struct {
	Date     time.Time `json:"date"`
	Title    string    `json:"title"`
	Location Location  `json:"location"`
}

func FromODPTBusStopPole(bsp []*odpt.BusStopPole) []*BusStopPole {
	esbsp := make([]*BusStopPole, len(bsp))[:0]

	for _, entry := range bsp {
		esbsp = append(esbsp, &BusStopPole{
			Date:  entry.Date.Date,
			Title: entry.Title,
			Location: Location{
				Latitude:  entry.Latitude,
				Longitude: entry.Longitude,
			},
		})
	}

	return esbsp
}
