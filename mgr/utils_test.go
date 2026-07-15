package mgr

import "testing"

func TestDetermineTempCharacteristic(t *testing.T) {
	tests := []struct {
		name      string
		isCelsius bool
		temp      float64
		expected  string
	}{
		{name: "hot fahrenheit", isCelsius: false, temp: 80, expected: "Hot"},
		{name: "cold fahrenheit", isCelsius: false, temp: 40, expected: "Cold"},
		{name: "moderate fahrenheit", isCelsius: false, temp: 60, expected: "Moderate"},
		{name: "boundary hot fahrenheit", isCelsius: false, temp: 76, expected: "Hot"},
		{name: "boundary cold fahrenheit", isCelsius: false, temp: 49, expected: "Cold"},
		{name: "boundary moderate fahrenheit", isCelsius: false, temp: 50, expected: "Moderate"},
		{name: "boundary moderate upper fahrenheit", isCelsius: false, temp: 75, expected: "Moderate"},
		{name: "hot celsius", isCelsius: true, temp: 25, expected: "Hot"},
		{name: "cold celsius", isCelsius: true, temp: 9, expected: "Cold"},
		{name: "moderate celsius", isCelsius: true, temp: 15, expected: "Moderate"},
		{name: "boundary hot celsius", isCelsius: true, temp: 24, expected: "Moderate"},
		{name: "boundary cold celsius", isCelsius: true, temp: 10, expected: "Moderate"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := determineTempCharacteristic(tt.isCelsius, tt.temp)
			if got != tt.expected {
				t.Fatalf("determineTempCharacteristic(%t, %v) = %q, want %q", tt.isCelsius, tt.temp, got, tt.expected)
			}
		})
	}
}
