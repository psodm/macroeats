package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/psodm/macroeats/internal/data"

	_ "github.com/lib/pq"
)

type FoodStore struct {
	DB *sql.DB
}

func NewFoodStore(db *sql.DB) *FoodStore {
	return &FoodStore{
		DB: db,
	}
}

func (f *FoodStore) Insert(ctx context.Context, food *data.Food) error {
	query := `INSERT INTO foods (food_name, brand_id, serving_size, serving_measurement_id, macros_id)
	          VALUES($1, $2, $3, $4, $5) RETURNING food_id`
	args := []any{food.FoodName, food.BrandID, food.ServingSize, food.ServingMeasurementID, food.MacrosID}
	err := f.DB.QueryRowContext(ctx, query, args...).Scan(&food.ID)
	if err != nil {
		switch {
		case data.IsDuplicate(err):
			return data.ErrDuplicateRow
		default:
			return err
		}
	}
	return nil
}

func (f *FoodStore) InsertTx(
	ctx context.Context,
	food *data.FoodTx,
	brandStore *BrandStore,
	measurementStore *MeasurementStore,
	macrosStore *MacrosStore) error {
	query := `INSERT INTO foods (food_name, brand_id, serving_size, serving_measurement_id, macros_id)
		VALUES($1, $2, $3, $4, $5) RETURNING food_id`
	tx, err := f.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	return transaction(tx, func() error {
		measurement, err := measurementStore.GetByAbbreviation(ctx, food.Measurement.MeasurementAbbreviation)
		if err != nil {
			return err
		} else {
			food.Measurement = *measurement
		}
		brand, err := brandStore.GetByName(ctx, food.Brand.BrandName)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				id, err := brandStore.InsertTx(ctx, tx, &food.Brand)
				if err != nil {
					return err
				}
				food.Brand.ID = id
			default:
				return err
			}
		} else {
			food.Brand = *brand
		}
		macrosID, err := macrosStore.InsertTx(ctx, tx, &food.Macros)
		if err != nil {
			return err
		}
		args := []any{
			food.FoodName,
			food.Brand.ID,
			food.ServingSize,
			food.Measurement.ID,
			macrosID,
		}
		err = tx.QueryRowContext(ctx, query, args...).Scan(&food.ID)
		if err != nil {
			return err
		}
		food.Macros.ID = macrosID
		return nil
	})
}
