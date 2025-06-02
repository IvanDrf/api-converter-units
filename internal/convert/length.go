package convert

import (
	"errors"

	"github.com/IvanDrf/units/internal/models"
)

const (
	millimeter = "mm"
	centimetre = "cm"
	meter      = "m"
	kilometer  = "km"

	inch = "in"
	feet = "ft"
	yard = "yd"
	mile = "mi"
)

func ConvertLength(req *models.Request) (float64, error) {
	switch req.Units {
	case millimeter:
		return FromMillimeters(req.Value, req.NewUnits)

	case centimetre:
		return FromCentimetre(req.Value, req.NewUnits)

	case meter:
		return FromMeters(req.Value, req.NewUnits)

	case kilometer:
		return FromKilometers(req.Value, req.NewUnits)

	case inch:
		return FromInches(req.Value, req.NewUnits)

	case feet:
		return FromFeets(req.Value, req.NewUnits)

	case yard:
		return FromYards(req.Value, req.NewUnits)

	case mile:
		return FromMiles(req.Value, req.NewUnits)

	default:
		return -1, errors.New("invalid units for length converting")
	}
}

func FromMeters(value float64, units string) (float64, error) {
	switch units {
	case millimeter:
		return value * 1000, nil

	case centimetre:
		return value * 100, nil

	case meter:
		return value, nil

	case kilometer:
		return value / 1000, nil

	case inch:
		return value * 39.370, nil

	case feet:
		return value * 3.28, nil

	case yard:
		return value * 1.0936, nil

	case mile:
		return value / 1609.344, nil

	default:
		return -1, errors.New("invalid units for converting from meters")

	}
}

func FromMillimeters(value float64, units string) (float64, error) {
	return FromMeters(value/100, units)
}

func FromCentimetre(value float64, units string) (float64, error) {
	return FromMeters(value/100, units)
}

func FromKilometers(value float64, units string) (float64, error) {
	return FromMeters(value*1000, units)
}

func FromInches(value float64, units string) (float64, error) {
	return FromMeters(value/39.370, units)
}

func FromFeets(value float64, units string) (float64, error) {
	return FromMeters(value/3.28, units)
}

func FromYards(value float64, units string) (float64, error) {
	return FromMeters(value/1.0936, units)
}

func FromMiles(value float64, units string) (float64, error) {
	return FromMeters(value*1609.344, units)
}
