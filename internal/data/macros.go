package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/psodm/macroeats/internal/validator"
)

const KJ_TO_CAL = 0.239006

type Macros struct {
	ID            int64   `kson:"Id"`
	Energy        float64 `json:"energy"`
	Calories      float64 `json:"calories"`
	Protein       float64 `json:"protein"`
	Carbohydrates float64 `json:"carbohydrates"`
	Fat           float64 `json:"fat"`
}

func ValidateMacros(v *validator.Validator, macros Macros) {
	v.Check(macros.Energy >= 0 && macros.Calories >= 0, "energy or calories", "must be provided and a positive number")
	v.Check(macros.Protein >= 0, "protein", "must be a positive number")
	v.Check(macros.Carbohydrates >= 0, "carbohydrates", "must be a positive number")
	v.Check(macros.Fat >= 0, "fat", "must be a positive number")
}

func (m *Macros) NormaliseMacrosEnergyAndCalories() {
	if m.Energy > 0 {
		m.Calories = m.Energy * KJ_TO_CAL
	} else {
		m.Energy = m.Calories / KJ_TO_CAL
	}
}

type MacrosModel struct {
	DB *sql.DB
}

func (m MacrosModel) Insert(macros *Macros) error {
	query := `INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
	          VALUES($1, $2, $3, $4, $5) RETURNING macros_id`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	args := []any{macros.Energy, macros.Calories, macros.Protein, macros.Carbohydrates, macros.Fat}
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&macros.ID)
}
