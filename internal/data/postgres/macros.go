package postgres

import (
	"context"
	"database/sql"

	"github.com/psodm/macroeats/internal/data"

	_ "github.com/lib/pq"
)

type MacrosStore struct {
	DB *sql.DB
}

func NewMacrosStore(db *sql.DB) *MacrosStore {
	return &MacrosStore{
		DB: db,
	}
}

func (m *MacrosStore) Insert(ctx context.Context, macros *data.Macros) error {
	query := `INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
	        VALUES($1, $2, $3, $4, $5) RETURNING macros_id`
	args := []any{macros.Energy, macros.Calories, macros.Protein, macros.Carbohydrates, macros.Fat}
	row := m.DB.QueryRowContext(ctx, query, args...)
	return row.Scan(&macros.ID)
}

func (m *MacrosStore) InsertTx(ctx context.Context, tx *sql.Tx, macros *data.Macros) (int64, error) {
	query := `INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
	        VALUES($1, $2, $3, $4, $5) RETURNING macros_id`
	args := []any{macros.Energy, macros.Calories, macros.Protein, macros.Carbohydrates, macros.Fat}
	var id int64
	row := tx.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
