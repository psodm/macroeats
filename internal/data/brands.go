package data

import (
	"context"
	"database/sql"
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
