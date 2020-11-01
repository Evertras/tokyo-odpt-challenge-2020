package esdata

import (
	"time"

	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

type Bus struct {
	Date                        time.Time `json:"date"`
	Location                    Location  `json:"location"`
	Operator                    string    `json:"operator"`
	Route                       string    `json:"route"`
	Number                      string    `json:"busNumber"`
	StartingBusStopPole         string    `json:"startingPole"`
	StartingBusStopPoleLocation Location  `json:"startingPoleLocation"`
	TerminalBusStopPole         string    `json:"terminalPole"`
	TerminalBusStopPoleLocation Location  `json:"terminalPoleLocation"`
	FromBusStop                 string    `json:"fromStop"`
	FromBusStopLocation         Location  `json:"fromStopLocation"`
	FromBusStopTime             time.Time `json:"leftLastStopAt"`
	ToBusStop                   string    `json:"toStop,omitempty"`
	ToBusStopLocation           *Location `json:"toStopLocation,omitempty"`
	ProgressPercent             int       `json:"progressPercent"`
	SpeedKmPerHour              float32   `json:"speedKmPerHour"`
	FacingDegrees               float32   `json:"azimuth"`
	DoorStatus                  string    `json:"doorStatus"`
}

func FromODPTBus(b []*odpt.Bus, poles odpt.BusStopPoleLookup) []*Bus {

	esb := make([]*Bus, len(b))

	for i, bus := range b {
		startingPoleLoc := poles[bus.StartingBusStopPole].Location
		terminalPoleLoc := poles[bus.TerminalBusStopPole].Location
		fromPole := poles[bus.FromBusStopPole]
		fromPoleLoc := odpt.Location{}
		if fromPole != nil {
			fromPoleLoc = poles[bus.FromBusStopPole].Location
		}

		busLoc := Location{}

		if bus.Latitude != 0 {
			busLoc = FromODPTLocation(bus.Location)
		} else {
			busLoc = FromODPTLocation(fromPoleLoc)
		}

		converted := &Bus{
			Date:                        bus.Date,
			Location:                    busLoc,
			Operator:                    removeType(bus.Operator),
			Route:                       removeFirstPeriodSeparatedChunk(removeType(bus.Route)),
			Number:                      removeType(bus.Number),
			StartingBusStopPole:         removeFirstPeriodSeparatedChunk(removeType(bus.StartingBusStopPole)),
			StartingBusStopPoleLocation: FromODPTLocation(startingPoleLoc),
			TerminalBusStopPole:         removeFirstPeriodSeparatedChunk(removeType(bus.TerminalBusStopPole)),
			TerminalBusStopPoleLocation: FromODPTLocation(terminalPoleLoc),
			FromBusStop:                 removeFirstPeriodSeparatedChunk(removeType(bus.FromBusStopPole)),
			FromBusStopLocation:         FromODPTLocation(fromPoleLoc),
			ProgressPercent:             int(bus.ProgressPercent0to1 * 100.),
			SpeedKmPerHour:              bus.SpeedKmPerHour,
			FacingDegrees:               bus.FacingDegrees,
			DoorStatus:                  bus.DoorStatus,
		}

		if bus.ToBusStopPole != "" {
			converted.ToBusStop = removeType(bus.ToBusStopPole)
			toLoc := FromODPTLocation(poles[bus.ToBusStopPole].Location)
			converted.ToBusStopLocation = &toLoc
		}

		esb[i] = converted
	}

	return esb
}
