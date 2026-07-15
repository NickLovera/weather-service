package models

type PointMetaData struct {
	Properties PointProperties `json:"properties"`
}

type PointProperties struct {
	//ForecastHourlyURL string   `json:"forecastHourly"` //opted to pull values out so repo function could be used "properly"
	RelativeLocation Location `json:"relativeLocation"`
	GridId           string   `json:"gridId"`
	GridX            int      `json:"gridX"`
	GridY            int      `json:"gridY"`
}

type Location struct {
	Properties SubProps `json:"properties"`
}

type SubProps struct {
	City  string `json:"city"`
	State string `json:"state"`
}
