package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/psodm/macroeats/internal/data"

	_ "github.com/lib/pq"
)

type RecipeStore struct {
	DB *sql.DB
}

func NewRecipeStore(db *sql.DB) *RecipeStore {
	return &RecipeStore{
		DB: db,
	}
}

func (r *RecipeStore) Insert(ctx context.Context, recipe *data.Recipe) error {
	query := `INSERT INTO recipes (recipe_name, recipe_description, recipe_meal_id, servings, prep_time, total_time, recipe_macros_id)
	          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING (recipe_id, created_at, version)`
	args := []any{recipe.Name, recipe.Description, recipe.MealID, recipe.Servings, recipe.PrepTime, recipe.TotalTime, recipe.MacrosID}
	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&recipe.ID, &recipe.CreatedAt, &recipe.Version)
	if err != nil {
		switch {
		case data.IsDuplicate(err):
			return data.ErrDuplicateRow
		default:
			return err
		}
	}
	return nil
}

func (r *RecipeStore) InsertTx(
	ctx context.Context,
	recipe *data.RecipeTx,
	ingredientSectionStore *IngredientSectionStore,
	ingredientStore *IngredientStore,
	foodStore *FoodStore,
	instructionStore *InstructionStore,
	noteStore *NoteStore,
	macrosStore *MacrosStore) error {
	query := `INSERT INTO recipes (recipe_name, recipe_description, recipe_meal_id, servings, prep_time, total_time, recipe_macros_id)
	          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING (recipe_id, created_at, version)`
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	return transaction(tx, func() error {
		fmt.Println(query)
		return nil
	})
}
