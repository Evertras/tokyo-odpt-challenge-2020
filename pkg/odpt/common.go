package odpt

import (
	"time"

	"github.com/dhconnelly/rtreego"
)

const (
	OpeningDoorFront = "odpt:OpeningDoor:FrontSide"
	OpeningDoorRear  = "odpt:OpeningDoor:RearSide"
)

type Base struct {
	ContextURL string    `json:"@context"`
	ID         string    `json:"@id"`
	Type       string    `json:"@type"`
	SameAs     string    `json:"owl:sameAs"`
	Date       time.Time `json:"dc:date"`
}

type Location struct {
	Longitude float64 `json:"geo:long"`
	Latitude  float64 `json:"geo:lat"`
}

func (l Location) Bounds() *rtreego.Rect {
	var point rtreego.Point = []float64{l.Latitude, l.Longitude}
	return point.ToRect(0.00000001)
}

type Valid struct {
	Valid time.Time `json:"dct:valid"`
}
