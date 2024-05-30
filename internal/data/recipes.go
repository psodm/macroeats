package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/psodm/macroeats/internal/validator"
)

type Recipe struct {
	ID          int64     `json:"recipeId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MealID      int64     `json:"mealId"`
	Servings    float64   `json:"servings"`
	PrepTime    int64     `json:"prepTime"`
	TotalTime   int64     `json:"totalTime"`
	MacrosID    int64     `json:"macrosId"`
	CreatedAt   time.Time `json:"-"`
	Version     int64     `json:"version"`
}

type RecipeTx struct {
	ID           int64               `json:"id"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Meal         Meal                `json:"meal"`
	Cuisines     []string            `json:"cuisines"`
	Servings     float64             `json:"servings"`
	PrepTime     int64               `json:"prepTime"`
	TotalTime    int64               `json:"totalTime"`
	Macros       Macros              `json:"macros"`
	Ingredients  map[string][]string `json:"ingredients"`
	Instructions []string            `json:"instructions"`
	Notes        []string            `json:"notes"`
	CreatedAt    time.Time           `json:"-"`
	Version      int64               `json:"version"`
}

func ValidateRecipe(v *validator.Validator, recipe *RecipeTx) {
	v.Check(recipe.Name != "", "name", "must be provided")
	v.Check(len(recipe.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(recipe.Description != "", "description", "must be provided")
	v.Check(len(recipe.Description) <= 1024, "description", "must not be more than 1024 bytes long")
	v.Check(recipe.Ingredients != nil, "ingredients", "must be provided")
	v.Check(len(recipe.Ingredients) >= 1, "ingredients", "must contain at least 1 ingredient")
	v.Check(recipe.Instructions != nil, "instructions", "must be provided")
	v.Check(len(recipe.Instructions) >= 1, "instructions", "must contain at least 1 instruction")
	v.Check(validator.Unique(recipe.Instructions), "instructions", "must not contain duplicate instructions")

	if recipe.Servings != 0 {
		v.Check(recipe.Servings > 0, "servings", "must be a positive number")
	}

	if recipe.PrepTime != 0 {
		v.Check(recipe.PrepTime > 0, "prepTime", "must be a positive number")
	}

	if recipe.TotalTime != 0 {
		v.Check(recipe.TotalTime > 0, "totalTime", "must be a positve number")
		if recipe.PrepTime > 0 {
			v.Check(recipe.TotalTime >= recipe.PrepTime, "totalTime", "must be equal to or greater than prepTime")
		}
	}
}

type RecipeModel struct {
	DB *sql.DB
}

func (r RecipeModel) Insert(recipe *Recipe) error {
	query := `INSERT INTO recipes (recipe_name, recipe_description, recipe_meal_id, servings, prep_time, total_time, recipe_macros_id, created_at, version)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING (recipe_id)`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	args := []any{recipe.Name, recipe.Description, recipe.MealID, recipe.Servings, recipe.PrepTime, recipe.TotalTime, recipe.MacrosID, recipe.CreatedAt}
	return r.DB.QueryRowContext(ctx, query, args...).Scan(&recipe.ID, &recipe.Version)
}

func (r RecipeModel) InsertTx(recipeTx *RecipeTx) error {
	fail := func(err error) error {
		return fmt.Errorf("CreateRecipe: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Begin transaction
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()
	query1 := `SELECT meal_id FROM meals WHERE meal_name = $1`
	query2 := `INSERT INTO macros (energy, calories, protein, carbohydrate, fat) VALUES ($1, $2, $3, $4, $5) RETURNING macros_id`
	query3 := `SELECT cuisine_id, cuisine_name from cuisines`
	query4 := `INSERT INTO recipe_cuisines (recipe_id, cuisine_id) VALUES($1, $2)`
	query5 := `INSERT INTO recipe_instructions(recipe_id, step, instruction) VALUES ($1, $2, $3) RETURNING instruction_id`
	query6 := `INSERT INTO recipe_notes(recipe_id, note_text)`
	// query7 := `INSERT INTO recipe_ingredient_sections (section_name) VALUES ($1) RETURNING section_id`

	err = tx.QueryRowContext(ctx, query1, recipeTx.Meal.MealName).Scan(&recipeTx.Meal.ID)
	if err != nil {
		return fail(err)
	}
	args := []any{recipeTx.Macros.Energy, recipeTx.Macros.Calories, recipeTx.Macros.Protein, recipeTx.Macros.Carbohydrates, recipeTx.Macros.Fat}
	err = tx.QueryRowContext(ctx, query2, args...).Scan(&recipeTx.Macros.ID)
	if err != nil {
		switch {
		case isDuplicate(err):
			return fail(ErrDuplicateRow)
		default:
			return fail(err)
		}
	}
	recipe := Recipe{0, recipeTx.Name, recipeTx.Description, recipeTx.Meal.ID, recipeTx.Servings, recipeTx.PrepTime, recipeTx.TotalTime, recipeTx.Macros.ID, recipeTx.CreatedAt, recipeTx.Version}
	err = r.Insert(&recipe)
	if err != nil {
		return fail(err)
	}
	recipeTx.ID = recipe.ID
	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	var cuisines []Cuisine

	// CONTINE HERE
	// cuisines
	rows, err := tx.QueryContext(ctx, query3, nil)
	if err != nil {
		return fail(err)
	}
	for rows.Next() {
		var cuisine Cuisine
		err = rows.Scan(&cuisine.ID, &cuisine.CuisineName)
		if err != nil {
			return fail(err)
		}
		cuisines = append(cuisines, cuisine)
	}
	for _, cuisine := range recipeTx.Cuisines {
		cuisineId := getCuisine(cuisines, cuisine)
		if cuisineId == -1 {
			return fail(fmt.Errorf("cuisine '%s' does not exist", cuisine))
		}
		err = tx.QueryRowContext(ctx, query4, recipeTx.ID, cuisineId).Scan()
		if err != nil {
			return fail(err)
		}
	}

	// ingredient sections
	// ingredients

	// for key, values := range recipeTx.Ingredients {
	// 	section := RecipeIngredientSection{}
	// 	err = tx.QueryRowContext(ctx, query7, key).Scan(&section.ID)
	// 	if err != nil {
	// 		return fail(err)
	// 	}
	// 	for _, ingredient := range values {
	// 		args = []any{recipeTx.ID, } //NO. Need to search for the food and include, or add
	// 	}
	// }

	// instructions
	for idx, value := range recipeTx.Instructions {
		instruction := RecipeInstruction{0, recipeTx.ID, int64(idx + 1), value}
		err = tx.QueryRowContext(ctx, query5, args...).Scan(&instruction.ID)
		if err != nil {
			return fail(err)
		}
	}

	// notes
	for _, note := range recipeTx.Notes {
		args := []any{recipeTx.ID, note}
		err = tx.QueryRowContext(ctx, query6, args...).Scan()
		if err != nil {
			return fail(err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	// End transaction
	return nil
}

func getCuisine(cuisines []Cuisine, cuisine string) int64 {
	for _, c := range cuisines {
		if cuisine == c.CuisineName {
			return c.ID
		}
	}
	return -1
}
