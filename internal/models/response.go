package models

type Responce struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UnitsType string  `json:"type"`
	Units     string  `json:"units"`
	Value     float64 `json:"value"`

	NewUnits string  `json:"converted_units"`
	NewValue float64 `json:"converted_value"`
}
