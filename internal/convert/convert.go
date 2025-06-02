package convert

import (
	"errors"

	"github.com/IvanDrf/units/internal/models"
)

const (
	length      = "length"
	temperature = "temperature"
	weight      = "weight"
)

func Convert(req *models.Request) (float64, error) {
	switch req.UnitsType {
	case length:
		return ConvertLength(req)

	case temperature:
		return ConvertTemperature(req)

	case weight:
		return ConvertWeight(req)

	default:
		return -1, errors.New("invalid units for converting")

	}

}
