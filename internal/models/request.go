package models

type Request struct {
	Units string  `json:"units"`
	Value float64 `json:"value"`

	NewUnits string `json:"converted_units"`
}
