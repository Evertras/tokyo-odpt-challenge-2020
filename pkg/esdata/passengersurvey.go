package esdata

import (
	"strings"
	"time"

	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

type PassengerSurvey struct {
	Date             time.Time `json:"date"`
	IncludeAlighting bool      `json:"includeAlighting"`
	Operator         string    `json:"operator"`
	SurveyYear       int       `json:"surveyYear"`
	PassengersPerDay int       `json:"passengersPerDay"`
	Line             string    `json:"line"`
	Station          string    `json:"station"`
	Location         Location  `json:"location"`
}

func FromODPTPassengerSurvey(ps []*odpt.PassengerSurvey, stations odpt.StationLookup) []*PassengerSurvey {
	esps := make([]*PassengerSurvey, len(ps))[:0]

	for _, entry := range ps {
		for _, station := range entry.Station {
			stationData := stations[station]

			if stationData == nil {
				panic(station)
			}

			for _, railway := range entry.Railway {
				trimmed := removeType(railway)
				split := strings.SplitN(trimmed, ".", 2)

				line := split[0]

				if len(split) > 1 {
					line = split[1]
				}

				for _, surveyObject := range entry.PassengerSurveyObjects {
					esps = append(esps, &PassengerSurvey{
						Date:             entry.Date.Date,
						IncludeAlighting: entry.IncludeAlighting,
						Operator:         removeType(entry.Operator),
						SurveyYear:       surveyObject.SurveyYear,
						PassengersPerDay: surveyObject.PassengerJourneys,
						Line:             line,
						Station:          removeType(stationData.Title),

						Location: Location{
							Latitude:  stationData.Latitude,
							Longitude: stationData.Longitude,
						},
					})
				}
			}
		}
	}

	return esps
}
