package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/psodm/macroeats/internal/validator"
)

type Food struct {
	ID                   int64   `json:"Id"`
	FoodName             string  `json:"foodName"`
	BrandID              int64   `json:"brandId"`
	ServingSize          float64 `json:"servingSize"`
	ServingMeasurementID int64   `json:"servingMeasurementId"`
	MacrosID             int64   `json:"macrosId"`
}

type FoodTx struct {
	ID          int64       `json:"Id"`
	FoodName    string      `json:"foodName"`
	Brand       Brand       `json:"brand"`
	ServingSize float64     `json:"servingSize"`
	Measurement Measurement `json:"measurement"`
	Macros      Macros      `json:"macros"`
}

func ValidateFood(v *validator.Validator, food Food) {
	v.Check(food.FoodName != "", "foodName", "must be provided")
	v.Check(food.ServingSize > 0, "servingSize", "must be provided")
}

type FoodModel struct {
	DB *sql.DB
}

func (f FoodModel) Insert(food *Food) error {
	query := `INSERT INTO foods (food_name, brand_id, serving_size, serving_measurement_id, macros_id)
	          VALUES($1, $2, $3, $4, $5) RETURNING food_id`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	args := []any{food.FoodName, food.BrandID, food.ServingSize, food.ServingMeasurementID, food.MacrosID}
	return f.DB.QueryRowContext(ctx, query, args...).Scan(&food.ID)
}

func (f FoodModel) InsertTx(foodTx *FoodTx) error {
	fail := func(err error) error {
		return fmt.Errorf("CreateFood: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Begin transaction
	tx, err := f.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()
	query1 := `SELECT measurement_id, measurement_name FROM measurements WHERE measurement_abbreviation = $1`
	query2 := `SELECT brand_id FROM brands WHERE brand_name = $1`
	query3 := `INSERT INTO brands (brand_name) VALUES ($1) RETURNING brand_id`
	query4 := `INSERT INTO macros (energy, calories, protein, carbohydrate, fat) VALUES ($1, $2, $3, $4, $5) RETURNING macros_id`
	err = tx.QueryRowContext(ctx, query1, foodTx.Measurement.MeasurementAbbreviation).Scan(&foodTx.Measurement.ID, &foodTx.Measurement.MeasurementName)
	if err != nil {
		return fail(err)
	}
	err = tx.QueryRowContext(ctx, query2, foodTx.Brand.BrandName).Scan(&foodTx.Brand.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			err = tx.QueryRowContext(ctx, query3, foodTx.Brand.BrandName).Scan(&foodTx.Brand.ID)
			if err != nil {
				return fail(err)
			}
		default:
			return fail(err)
		}
	}
	args := []any{foodTx.Macros.Energy, foodTx.Macros.Calories, foodTx.Macros.Protein, foodTx.Macros.Carbohydrates, foodTx.Macros.Fat}
	err = tx.QueryRowContext(ctx, query4, args...).Scan(&foodTx.Macros.ID)
	if err != nil {
		return fail(err)
	}
	food := Food{0, foodTx.FoodName, foodTx.Brand.ID, foodTx.ServingSize, foodTx.Measurement.ID, foodTx.Macros.ID}
	err = f.Insert(&food)
	if err != nil {
		return fail(err)
	}
	foodTx.ID = food.ID
	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	// End transaction
	return nil
}

func (f FoodModel) GetFoodById(id int64) (*Food, error) {
	return nil, nil
}
