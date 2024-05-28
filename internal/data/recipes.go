package data

import (
	"time"

	"github.com/psodm/macroeats/internal/validator"
)

type Recipe struct {
	ID           int64                     `json:"recipeId"`
	Name         string                    `json:"name"`
	Description  string                    `json:"description"`
	MealType     string                    `json:"mealType"`
	Servings     float64                   `json:"servings"`
	Macros       Macros                    `json:"macros"`
	PrepTime     CookingTime               `json:"prepTime"`
	TotalTime    CookingTime               `json:"totalTime"`
	Ingredients  []RecipeIngredientSection `json:"ingredients"`
	Instructions []RecipeInstruction       `json:"instructions"`
	Notes        []string                  `json:"notes"`
	CreatedAt    time.Time                 `json:"-"`
	Version      int64                     `json:"version"`
}

func ValidateRecipe(v *validator.Validator, recipe *Recipe) {
	v.Check(recipe.Name != "", "name", "must be provided")
	v.Check(len(recipe.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(recipe.Description != "", "description", "must be provided")
	v.Check(len(recipe.Description) <= 1024, "description", "must not be more than 1024 bytes long")
	// v.Check(recipe.Macros.Energy > 0, "macros energy", "must be provided and must be a positive number")
	v.Check(recipe.Ingredients != nil, "ingredients", "must be provided")
	v.Check(len(recipe.Ingredients) >= 1, "ingredients", "must contain at least 1 ingredient")
	// v.Check(validator.Unique(recipe.Ingredients), "ingredients", "must not contain duplicate values")
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
