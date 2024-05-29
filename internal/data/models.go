package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Brands       BrandModel
	Foods        FoodModel
	Macros       MacrosModel
	Measurements MeasurementModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Brands:       BrandModel{DB: db},
		Foods:        FoodModel{DB: db},
		Macros:       MacrosModel{DB: db},
		Measurements: MeasurementModel{DB: db},
	}
}
