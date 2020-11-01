package odpt

type BusStopPoleOrder struct {
	Pole                 string   `json:"odpt:busstopPole"`
	Index                int      `json:"odpt:index"`
	OpeningDoorsToGetOn  []string `json:"odpt:openingDoorsToGetOn"`
	OpeningDoorsToGetOff []string `json:"odpt:openingDoorsToGetOff"`
	Note                 string   `json:"odpt:Note"`
}
