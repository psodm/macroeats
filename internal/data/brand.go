package data

import (
	"github.com/psodm/macroeats/internal/validator"
)

type Brand struct {
	ID        int64  `json:"id"`
	BrandName string `json:"name"`
}

func ValidateBrand(v *validator.Validator, brand Brand) {
	// v.Check(brand.BrandName != "", "brandName", "must be provided")
}
