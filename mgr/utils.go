package mgr

func determineTempCharacteristic(isCelsius bool, temp float64) string {
	if isCelsius {
		if temp > 24 { // 75°F in Celsius
			return "Hot"
		} else if temp < 10 { // 50°F in Celsius
			return "Cold"
		}
	} else {
		if temp > 75 {
			return "Hot"
		} else if temp < 50 {
			return "Cold"
		}
	}
	return "Moderate"
}
