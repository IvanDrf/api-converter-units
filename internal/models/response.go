package models

type Responce struct {
	UnitsType string  `json:"type" gorm:"primaryKey"`
	Units     string  `json:"units"`
	Value     float64 `json:"value"`

	NewUnits string  `json:"converted_units"`
	NewValue float64 `json:"converted_value"`
}
