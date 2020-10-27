package esdata

import (
	"strings"
	"time"

	"github.com/evertras/tokyo-odpt-challenge-2020/pkg/odpt"
)

type PassengerSurvey struct {
	Date              time.Time
	IncludeAlighting  bool
	Operator          string
	SurveyYear        int
	PassengerJourneys int
	Line              string
	Company           string
	Station           string
}

func FromODPTPassengerSurvey(ps []*odpt.PassengerSurvey) []*PassengerSurvey {
	esps := make([]*PassengerSurvey, len(ps))

	for i, entry := range ps {
		for _, surveyObject := range entry.PassengerSurveyObjects {
			for _, railway := range entry.Railway {
				trimmed := strings.Split(railway, ":")[1]
				split := strings.Split(trimmed, ".")

				esps[i] = &PassengerSurvey{
					Date:              entry.Date,
					IncludeAlighting:  entry.IncludeAlighting,
					Operator:          entry.Operator,
					SurveyYear:        surveyObject.SurveyYear,
					PassengerJourneys: surveyObject.PassengerJourneys,
					Line:              split[1],
					Station:           entry.Station[0],
				}
			}
		}
	}

	return esps
}
