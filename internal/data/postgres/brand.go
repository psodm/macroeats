package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/psodm/macroeats/internal/data"
)

type BrandStore struct {
	conn *pgx.Conn
}

func NewBrandStore(conn *pgx.Conn) *BrandStore {
	return &BrandStore{
		conn: conn,
	}
}

func (b *BrandStore) Insert(ctx context.Context, brand *data.Brand) error {
	sql := `INSERT INTO brands(brand_name) VALUES ($1) RETURNING id`
	row := b.conn.QueryRow(ctx, sql, &brand.BrandName)
	if err := row.Scan(&brand.ID); err != nil {
		return fmt.Errorf("insert brand: %w", err)
	}
	return nil
}

func (b *BrandStore) GetByName(ctx context.Context, name string) (data.Brand, error) {
	sql := `SELECT brand_id FROM brands WHERE brand_name = $1`
	row := b.conn.QueryRow(ctx, sql, &name)
	brand := data.Brand{}
	if err := row.Scan(&brand.ID); err != nil {
		return brand, fmt.Errorf("get brand: %w", err)
	}
	return brand, nil
}
