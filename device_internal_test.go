package keylight

import (
	"fmt"
	"testing"
)

func TestLightTemperatureConversion(t *testing.T) {
	tests := []struct {
		kelvin, elgato int
	}{
		{
			kelvin: 2900,
			elgato: 343,
		},
		{
			kelvin: 7000,
			elgato: 142,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%dK", tt.kelvin), func(t *testing.T) {
			kelvin := convertToKelvin(tt.elgato)
			if kelvin != tt.kelvin {
				t.Fatalf("unexpected temperature Kelvin value, expected %d but got %d", tt.kelvin, kelvin)
			}

			if elgato := convertToAPI(kelvin); elgato != tt.elgato {
				t.Fatalf("unexpected temperature Elgato value, expected %d but got %d", tt.elgato, elgato)
			}
		})
	}
}
