package models

type Responce struct {
	Units string  `json:"units"`
	Value float64 `json:"value"`

	NewUnits string  `json:"converted_units"`
	NewValue float64 `json:"converted_value"`
}
