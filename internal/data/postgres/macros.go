package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/psodm/macroeats/internal/data"
)

type MacrosStore struct {
	conn *pgx.Conn
}

func NewMacrosStore(conn *pgx.Conn) *MacrosStore {
	return &MacrosStore{
		conn: conn,
	}
}

func (m *MacrosStore) Insert(ctx context.Context, macros *data.Macros) error {
	sql := `INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
	        VALUES($1, $2, $3, $4, $5) RETURNING macros_id`
	args := []any{macros.Energy, macros.Calories, macros.Protein, macros.Carbohydrates, macros.Fat}
	row := m.conn.QueryRow(ctx, sql, args...)
	if err := row.Scan(&macros.ID); err != nil {
		return fmt.Errorf("insert macros: %w", err)
	}
	return nil
}
