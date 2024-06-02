package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/psodm/macroeats/internal/data"

	_ "github.com/lib/pq"
)

type BrandStore struct {
	DB *sql.DB
}

func NewBrandStore(db *sql.DB) *BrandStore {
	return &BrandStore{
		DB: db,
	}
}

func (b *BrandStore) Insert(ctx context.Context, brand *data.Brand) error {
	query := `INSERT INTO brands(brand_name) VALUES ($1) RETURNING brand_id`
	err := b.DB.QueryRowContext(ctx, query, &brand.BrandName).Scan(&brand.ID)
	if err != nil {
		return fmt.Errorf("insert brand: %w", err)
	}
	return nil
}

func (b *BrandStore) InsertTx(ctx context.Context, tx *sql.Tx, brand *data.Brand) (int64, error) {
	query := `INSERT INTO brands(brand_name) VALUES ($1) RETURNING brand_id`
	var id int64
	row := tx.QueryRowContext(ctx, query, brand.BrandName)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (b *BrandStore) Get(ctx context.Context, id int64) (*data.Brand, error) {
	query := `SELECT brand_id, brand_name FROM brands
			  WHERE brand_id = $1`
	var brand data.Brand
	err := b.DB.QueryRowContext(ctx, query, id).Scan(
		&brand.ID,
		&brand.BrandName,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, data.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &brand, nil
}

func (b *BrandStore) GetByName(ctx context.Context, name string) (*data.Brand, error) {
	var brand data.Brand
	query := `SELECT brand_id FROM brands WHERE brand_name = $1`
	err := b.DB.QueryRowContext(ctx, query, &name).Scan(&brand.ID)
	if err != nil {
		return nil, fmt.Errorf("get brand: %w", err)
	}
	brand.BrandName = name
	return &brand, nil
}

func (b *BrandStore) GetAll(ctx context.Context) ([]*data.Brand, error) {
	query := `SELECT brand_id, brand_name FROM brands`
	var brands []*data.Brand
	rows, err := b.DB.QueryContext(ctx, query)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, data.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	for rows.Next() {
		var brand data.Brand
		err := rows.Scan(
			&brand.ID,
			&brand.BrandName,
		)
		if err != nil {
			return nil, err
		}
		brands = append(brands, &brand)
	}
	return brands, nil
}

func (b *BrandStore) Update(ctx context.Context, brand *data.Brand) error {
	query := `UPDATE brands
	          SET brand_name = $1
			  WHERE brand_id = $2`
	args := []any{
		brand.BrandName,
		brand.ID,
	}
	return b.DB.QueryRowContext(ctx, query, args...).Scan()
}

func (b *BrandStore) Delete(id int64) error {
	return nil
}
