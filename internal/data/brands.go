package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/psodm/macroeats/internal/validator"
)

type Brand struct {
	ID        int64  `json:"Id"`
	BrandName string `json:"brandName"`
}

func ValidateBrand(v *validator.Validator, brand Brand) {
	// v.Check(brand.BrandName != "", "brandName", "must be provided")
}

type BrandModel struct {
	DB *sql.DB
}

func (b BrandModel) CreateBrand(brand *Brand) error {
	query := `INSERT INTO brands (brand_name) VALUES ($1) RETURNING brand_id`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return b.DB.QueryRowContext(ctx, query, brand.BrandName).Scan(&brand.ID)
}

func (b BrandModel) Get(id int64) (*Brand, error) {
	query := `SELECT brand_id, brand_name FROM brands
			  WHERE brand_id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var brand Brand
	err := b.DB.QueryRowContext(ctx, query, id).Scan(
		&brand.ID,
		&brand.BrandName,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &brand, nil
}

func (b BrandModel) GetByName(name string) (*Brand, error) {
	query := `SELECT brand_id, brand_name FROM brands
			  WHERE brand_name = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var brand Brand
	err := b.DB.QueryRowContext(ctx, query, name).Scan(
		&brand.ID,
		&brand.BrandName,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &brand, nil
}

func (b BrandModel) GetAll() ([]*Brand, error) {
	query := `SELECT brand_id, brand_name FROM brands`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var brands []*Brand
	rows, err := b.DB.QueryContext(ctx, query)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	for rows.Next() {
		var brand Brand
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

func (b BrandModel) UpdateBrand(brand *Brand) error {
	query := `UPDATE brands
	          SET brand_name = $1
			  WHERE brand_id = $2`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	args := []any{
		brand.BrandName,
		brand.ID,
	}
	return b.DB.QueryRowContext(ctx, query, args...).Scan()
}

func (b BrandModel) Delete(id int64) error {
	return nil
}
