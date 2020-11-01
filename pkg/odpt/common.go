package odpt

import "time"

const (
	OpeningDoorFront = "odpt:OpeningDoor:FrontSide"
	OpeningDoorRear  = "odpt:OpeningDoor:RearSide"
)

type Base struct {
	ContextURL string `json:"@context"`
	ID         string `json:"@id"`
	Type       string `json:"@type"`
	SameAs     string `json:"owl:sameAs"`
	Date time.Time `json:"dc:date"`
}

type Location struct {
	Longitude float64 `json:"geo:long"`
	Latitude  float64 `json:"geo:lat"`
}
