package esdata

import "github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func FromODPTLocation(loc odpt.Location) Location {
	return Location{
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
	}
}

// TODO: add nillable version maybe that returns nil on 0 vals?
