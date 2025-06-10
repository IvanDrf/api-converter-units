package convert

import (
	"errors"

	"github.com/IvanDrf/units/internal/models"
)

const (
	length      = "length"
	temperature = "temperature"
	weight      = "weight"
	time        = "time"
)

func Convert(req *models.Request) (float64, error) {
	switch req.UnitsType {
	case length:
		return ConvertLength(req)

	case temperature:
		return ConvertTemperature(req)

	case weight:
		return ConvertWeight(req)

	case time:
		return ConvertTime(req)

	default:
		return -1, errors.New("invalid units for converting")

	}

}

func CreateConversion(req *models.Request) (models.Responce, error) {
	result := models.Responce{
		UnitsType: req.UnitsType,
		Units:     req.Units,
		Value:     req.Value,

		NewUnits: req.NewUnits,
		NewValue: -1,
	}

	var err error
	result.NewValue, err = Convert(req)
	if err != nil {
		return models.Responce{}, err
	}

	return result, err
}
