package models

type CurrentWeather struct {
	City               string `json:"city"`
	State              string `json:"state"`
	Forecast           string `json:"forecast"`
	TempCharacteristic string `json:"tempCharacteristic"`
}
