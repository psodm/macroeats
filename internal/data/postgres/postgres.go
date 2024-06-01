package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func transaction(ctx context.Context, tx pgx.Tx, f func() error) error {
	if err := f(); err != nil {
		_ = tx.Rollback(ctx)
		return fmt.Errorf("f %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit %w", err)
	}
	return nil
}
