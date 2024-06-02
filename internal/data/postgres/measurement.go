package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/psodm/macroeats/internal/data"

	_ "github.com/lib/pq"
)

type MeasurementStore struct {
	DB *sql.DB
}

func NewMeansurementStore(db *sql.DB) *MeasurementStore {
	return &MeasurementStore{
		DB: db,
	}
}

func (m *MeasurementStore) Insert(ctx context.Context, measurement *data.Measurement) error {
	query := `INSERT INTO measurements (measurement_name, measurement_abbreviation)
	          VALUES ($1, $2) RETURNING measurement_id`
	args := []any{&measurement.MeasurementAbbreviation, &measurement.MeasurementAbbreviation}
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&measurement.ID)
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

func (m *MeasurementStore) Get(ctx context.Context, id int64) (*data.Measurement, error) {
	query := `SELECT measurement_id, measurement_name, measurement_abbreviation
	          FROM measurements
			  WHERE measurement_id = $1`
	var measurement data.Measurement
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&measurement.ID,
		&measurement.MeasurementName,
		&measurement.MeasurementAbbreviation,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, data.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &measurement, nil
}

func (b *MeasurementStore) GetByAbbreviation(ctx context.Context, abbreviation string) (*data.Measurement, error) {
	query := `SELECT measurement_id, measurement_name, measurement_abbreviation
	          FROM measurements
	          WHERE measurement_abbreviation = $1`
	var measurement data.Measurement
	err := b.DB.QueryRowContext(ctx, query, &abbreviation).Scan(
		&measurement.ID,
		&measurement.MeasurementName,
		&measurement.MeasurementAbbreviation,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, data.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &measurement, nil
}

func (m *MeasurementStore) GetAll(ctx context.Context) ([]*data.Measurement, error) {
	query := `SELECT measurement_id, measurement_name, measurement_abbreviation
	          FROM measurements`
	var measurements []*data.Measurement
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, data.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	for rows.Next() {
		var measurement data.Measurement
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

// This is a very broad stroke approch. To improve
func (m *MeasurementStore) Update(ctx context.Context, measurement *data.Measurement) error {
	query := `UPDATE measurements
	          SET measurement_name = $1, measurement_abbreviation = $2
			  WHERE measurement_id = $3`
	args := []any{
		measurement.MeasurementName,
		measurement.MeasurementAbbreviation,
		measurement.ID,
	}
	return m.DB.QueryRowContext(ctx, query, args...).Scan()
}

func (m *MeasurementStore) Delete(id int64) error {
	return nil
}
