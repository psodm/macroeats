package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/psodm/macroeats/internal/validator"
)

type Food struct {
	ID                   int64   `json:"Id"`
	FoodName             string  `json:"foodName"`
	ServingQuantity      float64 `json:"servingSize"`
	ServingMeasurementID int64   `json:"servingMeasurementId"`
	MacrosID             int64   `json:"macrosId"`
}

type FoodTx struct {
	ID              int64       `json:"Id"`
	FoodName        string      `json:"foodName"`
	ServingQuantity float64     `json:"servingQuantity"`
	Measurement     Measurement `json:"measurement"`
	Macros          Macros      `json:"macros"`
}

func ValidateFood(v *validator.Validator, food Food) {
	v.Check(food.FoodName != "", "foodName", "must be provided")
	v.Check(food.ServingQuantity > 0, "servingSize", "must be provided")
}

type FoodModel struct {
	DB *sql.DB
}

func (f FoodModel) Insert(food *Food) error {
	query := `INSERT INTO foods (food_name, serving_quantity, serving_measurement_id, macros_id)
	          VALUES($1, $2, $3, $4) RETURNING food_id`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	args := []any{food.FoodName, food.ServingQuantity, food.ServingMeasurementID, food.MacrosID}
	return f.DB.QueryRowContext(ctx, query, args...).Scan(&food.ID)
}

func (f FoodModel) InsertTX(foodTx *FoodTx) error {
	fail := func(err error) error {
		return fmt.Errorf("CreateFood: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := f.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	query1 := `SELECT measurement_id, measurement_name FROM measurements WHERE measurement_abbreviation = $1`
	query2 := `INSERT INTO macros (energy, calories, protein, carbohydrate, fat) VALUES ($1, $2, $3, $4, $5) RETURNING macros_id`
	query3 := `INSERT into foods (food_name, serving_quantity, serving_measurement_id, macros_id) VALUES ($1, $2, $3, $4) RETURNING food_id`

	err = tx.QueryRowContext(ctx, query1, foodTx.Measurement.MeasurementAbbreviation).Scan(&foodTx.Measurement.ID, &foodTx.Measurement.MeasurementName)
	if err != nil {
		return fail(err)
	}

	args := []any{foodTx.Macros.Energy, foodTx.Macros.Calories, foodTx.Macros.Protein, foodTx.Macros.Carbohydrates, foodTx.Macros.Fat}
	err = tx.QueryRowContext(ctx, query2, args...).Scan(&foodTx.Macros.ID)
	if err != nil {
		return fail(err)
	}

	args = []any{foodTx.FoodName, foodTx.ServingQuantity, foodTx.Measurement.ID, foodTx.Measurement.ID}
	return tx.QueryRowContext(ctx, query3, args...).Scan(&foodTx.ID)
}
