package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/psodm/macroeats/internal/validator"
)

type Measurement struct {
	ID                      int64  `json:"id"`
	MeasurementName         string `json:"name"`
	MeasurementAbbreviation string `json:"abbreviation"`
}

func ValidateMeasurementUnit(v *validator.Validator, unit Measurement) {
	v.Check(unit.MeasurementName != "", "measurementName", "must be provided")
	v.Check(unit.MeasurementAbbreviation != "", "measurementAbbreviation", "must be provided")
}

type MeasurementModel struct {
	DB *sql.DB
}

func (m MeasurementModel) Insert(measurement *Measurement) error {
	query := `INSERT INTO measurements (measurement_name, measurement_abbreviation)
	          VALUES ($1, $2) RETURNING measurement_id`
	args := []any{measurement.MeasurementName, measurement.MeasurementAbbreviation}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&measurement.ID)
	if err != nil {
		switch {
		case isDuplicate(err):
			return ErrDuplicateRow
		default:
			return err
		}
	}
	return nil
}

func (m MeasurementModel) Get(id int64) (*Measurement, error) {
	query := `SELECT measurement_id, measurement_name, measurement_abbreviation
	          FROM measurements
			  WHERE measurement_id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var measurement Measurement
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&measurement.ID,
		&measurement.MeasurementName,
		&measurement.MeasurementAbbreviation,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &measurement, nil
}

func (m MeasurementModel) GetByAbbreviation(abbrev string) (*Measurement, error) {
	query := `SELECT measurement_id, measurement_name, measurement_abbreviation
	          FROM measurements
			  WHERE measurement_abbreviation = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var measurement Measurement
	err := m.DB.QueryRowContext(ctx, query, abbrev).Scan(
		&measurement.ID,
		&measurement.MeasurementName,
		&measurement.MeasurementAbbreviation,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &measurement, nil
}

func (m MeasurementModel) GetAll() ([]*Measurement, error) {
	query := `SELECT measurement_id, measurement_name, measurement_abbreviation
	          FROM measurements`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var measurements []*Measurement
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	for rows.Next() {
		var measurement Measurement
		err := rows.Scan(
			&measurement.ID,
			&measurement.MeasurementName,
			&measurement.MeasurementAbbreviation,
		)
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, &measurement)
	}
	return measurements, nil
}

func (m MeasurementModel) Update(measurement *Measurement) error {
	query := `UPDATE measurements
	          SET measurement_name = $1, measurement_abbreviation = $2
			  WHERE measurement_id = $3`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	args := []any{
		measurement.MeasurementName,
		measurement.MeasurementAbbreviation,
		measurement.ID,
	}
	return m.DB.QueryRowContext(ctx, query, args...).Scan()
}

func (m MeasurementModel) Delete(id int64) error {
	return nil
}
