package mgr

func determineTempCharacteristic(temp float64) string {
	if temp > 75 {
		return "Hot"
	} else if temp < 50 {
		return "Cold"
	}
	return "Moderate"
}
