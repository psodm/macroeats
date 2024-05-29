package data

import (
	"database/sql"
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
