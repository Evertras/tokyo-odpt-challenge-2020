package odpt

type Base struct {
	ContextURL             string                  `json:"@context"`
	ID                     string                  `json:"@id"`
	Type                   string                  `json:"@type"`
	SameAs     string    `json:"owl:sameAs"`
}

type Location struct {
	Longitude  float64   `json:"geo:long"`
	Latitude   float64   `json:"geo:lat"`
}
