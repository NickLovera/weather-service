package mgr

import "testing"

func TestDetermineTempCharacteristic(t *testing.T) {
	tests := []struct {
		name     string
		temp     float64
		expected string
	}{
		{name: "hot temperature", temp: 80, expected: "Hot"},
		{name: "cold temperature", temp: 40, expected: "Cold"},
		{name: "moderate temperature", temp: 60, expected: "Moderate"},
		{name: "boundary hot", temp: 76, expected: "Hot"},
		{name: "boundary cold", temp: 49, expected: "Cold"},
		{name: "boundary moderate", temp: 50, expected: "Moderate"},
		{name: "boundary moderate upper", temp: 75, expected: "Moderate"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := determineTempCharacteristic(tt.temp)
			if got != tt.expected {
				t.Fatalf("determineTempCharacteristic(%v) = %q, want %q", tt.temp, got, tt.expected)
			}
		})
	}
}
