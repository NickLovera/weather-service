package models

type GridPointHourly struct {
	Properties GridPointProperties `json:"properties"`
}

type GridPointProperties struct {
	Periods []Periods `json:"periods"`
}

type Periods struct {
	Temperature   interface{} `json:"temperature"`
	ShortForecast string      `json:"shortForecast"`
}
