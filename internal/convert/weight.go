package convert

import (
	"errors"

	"github.com/IvanDrf/units/internal/models"
)

const (
	milligram = "mg"
	gram      = "g"
	kilogram  = "kg"

	pound = "lb"
	ounce = "oz"
)

func ConvertWeight(req *models.Request) (float64, error) {
	switch req.Units {
	case milligram:
		return FromMilligram(req.Value, req.NewUnits)

	case gram:
		return FromGram(req.Value, req.NewUnits)

	case kilogram:
		return FromKilogram(req.Value, req.NewUnits)

	case pound:
		return FromPound(req.Value, req.NewUnits)

	case ounce:
		return FromOunce(req.Value, req.NewUnits)

	default:
		return -1, errors.New("invalid units for weight converting")

	}

}

func FromGram(value float64, units string) (float64, error) {
	switch units {
	case milligram:
		return value * 1000, nil

	case gram:
		return value, nil

	case kilogram:
		return value / 1000, nil

	case pound:
		return value / 453.59237, nil

	case ounce:
		return value / 28.34949, nil

	default:
		return -1, errors.New("invalid units for weight converting")

	}
}

func FromMilligram(value float64, units string) (float64, error) {
	return FromGram(value/1000, units)
}

func FromKilogram(value float64, units string) (float64, error) {
	return FromGram(value*1000, units)
}

func FromPound(value float64, units string) (float64, error) {
	return FromGram(value*453.59237, units)
}

func FromOunce(value float64, units string) (float64, error) {
	return FromGram(value*28.34949, units)
}
