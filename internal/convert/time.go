package convert

import (
	"errors"

	"github.com/IvanDrf/units/internal/models"
)

const (
	second = "s"
	minute = "min"

	hour = "h"
	day  = "day"
)

func ConvertTime(req *models.Request) (float64, error) {
	switch req.Units {
	case second:
		return FromSeconds(req.Value, req.NewUnits)

	case minute:
		return FromMinutes(req.Value, req.NewUnits)

	case hour:
		return FromHours(req.Value, req.NewUnits)

	case day:
		return FromDays(req.Value, req.NewUnits)

	default:
		return -1, errors.New("invalid units for time converting")
	}
}

func FromSeconds(value float64, units string) (float64, error) {
	switch units {
	case second:
		return value, nil

	case minute:
		return value / 60, nil

	case hour:
		return value / 3600, nil

	case day:
		return value / (3600 * 24), nil

	default:
		return -1, errors.New("invalid units for time converting")
	}
}

func FromMinutes(value float64, units string) (float64, error) {
	return FromSeconds(value*60, units)
}

func FromHours(value float64, units string) (float64, error) {
	return FromSeconds(value*3600, units)
}

func FromDays(value float64, units string) (float64, error) {
	return FromSeconds(value*24*3600, units)
}
