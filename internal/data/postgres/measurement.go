package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/psodm/macroeats/internal/data"
)

type MeasurementStore struct {
	conn *pgx.Conn
}

func NewMeansurementStore(conn *pgx.Conn) *MeasurementStore {
	return &MeasurementStore{
		conn: conn,
	}
}

func (m *MeasurementStore) Insert(ctx context.Context, measurement *data.Measurement) error {
	sql := `INSERT INTO measurements(measurement_name, measurement_abbreviation) VALUES ($1, $2) RETURNING id`
	row := m.conn.QueryRow(ctx, sql, &measurement.MeasurementAbbreviation, &measurement.MeasurementAbbreviation)
	if err := row.Scan(&measurement.ID); err != nil {
		return fmt.Errorf("insert measurement: %w", err)
	}
	return nil
}

func (b *MeasurementStore) GetByName(ctx context.Context, abbreviation string) (data.Measurement, error) {
	sql := `SELECT measurement_id, measurement_name, measurement_abbreviation FROM measurements WHERE measurement_abbreviation = $1`
	row := b.conn.QueryRow(ctx, sql, &abbreviation)
	measurement := data.Measurement{}
	if err := row.Scan(&measurement.ID, &measurement.MeasurementName, &measurement.MeasurementAbbreviation); err != nil {
		return measurement, fmt.Errorf("get measurement: %w", err)
	}
	return measurement, nil
}
