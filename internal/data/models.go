package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Measurements MeasurementModel
	Foods        FoodModel
	Macros       MacrosModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Measurements: MeasurementModel{DB: db},
		Foods:        FoodModel{DB: db},
		Macros:       MacrosModel{DB: db},
	}
}
