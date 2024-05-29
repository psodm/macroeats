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
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Meal         Meal      `json:"meal"`
	Servings     float64   `json:"servings"`
	PrepTime     int64     `json:"prepTime"`
	TotalTime    int64     `json:"totalTime"`
	Macros       Macros    `json:"macros"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	Notes        []string  `json:"notes"`
	CreatedAt    time.Time `json:"-"`
	Version      int64     `json:"version"`
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

	// CONTINE HERE
	// ingredients
	// instructions
	// notes

	// End transaction
	return nil
}
