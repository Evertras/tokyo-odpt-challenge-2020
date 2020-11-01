package odpt

import "time"

type Bus struct {
	Base
	Valid
	Location

	Number                 string    `json:"busNumber"`
	Operator               string    `json:"operator"`
	UpdateFrequencySeconds int       `json:"odpt:frequency"`
	StartingBusStopPole    string    `json:"odpt:startingBusstopPole"`
	TerminalBusStopPole    string    `json:"odpt:terminalBusstopPole"`
	FromBusStopPole        string    `json:"odpt:fromBusstopPole"`
	FromBusStopPoleTime    time.Time `json:"odpt:fromBusstopPoleTime"`
	ToBusStopPole          string    `json:"odpt:toBusstopPole"`
	ProgressPercent0to1    float32   `json:"odpt:progress"`
	SpeedKmPerHour         float32   `json:"odpt:speed"`
	FacingDegrees          float32   `json:"odpt:azimuth"`
	DoorStatus             string    `json:"odpt:doorStatus"`
}

