package convert

import (
	"errors"
  
	"github.com/IvanDrf/units/internal/models"
)

const (
	celsius    = "C"
	fahrenheit = "F"
	kelvin     = "K"
)

func ConvertTemperature(req *models.Request) (float64, error) {
	switch req.Units {
	case celsius:
		return FromCelsius(req.Value, req.NewUnits)

	case fahrenheit:
		return FromFahrenheit(req.Value, req.NewUnits)

	case kelvin:
		return FromKelvin(req.Value, req.NewUnits)

	default:
		return -1, errors.New("invalid units for temperature converting")
	}
}

func FromCelsius(value float64, units string) (float64, error) {
	switch units {
	case celsius:
		return value, nil

	case fahrenheit:
		return value*9/5 + 32, nil

	case kelvin:
		return value + 273.15, nil

	default:
		return -1, errors.New("invalid units for temperature converting")
	}
}

func FromFahrenheit(value float64, units string) (float64, error) {
	return FromCelsius((value-32)*5/9, units)
}

func FromKelvin(value float64, units string) (float64, error) {
	return FromCelsius(value-273.15, units)
}
