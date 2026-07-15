package models

type CurrentWeather struct {
	City               string  `json:"city"`
	State              string  `json:"state"`
	Forecast           string  `json:"forecast"`
	Temperature        float64 `json:"temperature"`
	TempCharacteristic string  `json:"tempCharacteristic"`
}
