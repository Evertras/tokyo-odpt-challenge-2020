package odpt

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type PassengerSurvey struct {
	ContextURL             string                  `json:"@context"`
	ID                     string                  `json:"@id"`
	Type                   string                  `json:"@type"`
	Date                   time.Time               `json:"dc:date"`
	SameAs                 string                  `json:"owl:sameAs"`
	Operator               string                  `json:"odpt:operator"`
	Station                []string                `json:"odpt:station"`
	Railway                []string                `json:"odpt:railway"`
	IncludeAlighting       bool                    `json:"odpt:includeAlighting"`
	PassengerSurveyObjects []PassengerSurveyObject `json:"odpt:passengerSurveyObject"`
}

type PassengerSurveyObject struct {
	SurveyYear        int `json:"odpt:surveyYear"`
	PassengerJourneys int `json:"odpt:passengerJourneys"`
}

// LoadPassengerSurveysJSON loads all PassengerSurvey entries from a static JSON
// file created by the data dump API
func LoadPassengerSurveysJSON(filename string) ([]*PassengerSurvey, error) {
	f, err := os.Open("./data/PassengerSurvey.json")

	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	ps := []*PassengerSurvey{}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&ps)

	if err != nil {
		return nil, fmt.Errorf("decoder.Decode: %w", err)
	}

	return ps, nil
}
