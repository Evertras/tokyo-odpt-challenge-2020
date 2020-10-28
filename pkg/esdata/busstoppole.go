package esdata

import (
	"time"

	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

type BusstopPole struct {
	Date     time.Time `json:"date"`
	Title    string    `json:"title"`
	Location Location  `json:"location"`
}

func FromODPTBusstopPole(bsp []*odpt.BusstopPole) []*BusstopPole {
	esbsp := make([]*BusstopPole, len(bsp))[:0]

	for _, entry := range bsp {
		esbsp = append(esbsp, &BusstopPole{
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
