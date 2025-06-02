package convert

import (
	"errors"
)

const (
	milligram = "mg"
	gram      = "g"
	kilogram  = "kg"

	pound = "lb"
	ounce = "oz"
)

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
		return -1, errors.New("invalid units for converting from weight")

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
