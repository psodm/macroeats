package data

import (
	"github.com/psodm/macroeats/internal/validator"
)

type Food struct {
	ID                   int64   `json:"id"`
	FoodName             string  `json:"name"`
	BrandID              int64   `json:"brandId"`
	ServingSize          float64 `json:"servingSize"`
	ServingMeasurementID int64   `json:"servingMeasurementId"`
	MacrosID             int64   `json:"macrosId"`
}

type FoodTx struct {
	ID          int64       `json:"id"`
	FoodName    string      `json:"name"`
	Brand       Brand       `json:"brand"`
	ServingSize float64     `json:"servingSize"`
	Measurement Measurement `json:"measurement"`
	Macros      Macros      `json:"macros"`
}

func ValidateFood(v *validator.Validator, food Food) {
	v.Check(food.FoodName != "", "foodName", "must be provided")
	v.Check(food.ServingSize > 0, "servingSize", "must be provided")
}
