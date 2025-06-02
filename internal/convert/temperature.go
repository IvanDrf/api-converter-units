package convert

import (
	"errors"
)

const (
	celsius    = "C"
	fahrenheit = "F"
	kelvin     = "K"
)

func FromCelsius(value float64, units string) (float64, error) {
	switch units {
	case celsius:
		return value, nil

	case fahrenheit:
		return value*9/5 + 32, nil

	case kelvin:
		return value + 273.15, nil

	default:
		return -1, errors.New("invalid units for converting from temperature")
	}
}

func FromFahrenheit(value float64, units string) (float64, error) {
	return FromCelsius((value-32)*5/9, units)
}

func FromKelvin(value float64, units string) (float64, error) {
	return FromCelsius(value-273.15, units)
}
