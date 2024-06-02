package postgres

import (
	"database/sql"
	"fmt"
)

func transaction(tx *sql.Tx, f func() error) error {
	if err := f(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("f %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit %w", err)
	}
	return nil
}
