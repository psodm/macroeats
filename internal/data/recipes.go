package data

import "time"

type Recipe struct {
	ID           int64                     `json:"recipeId"`
	Name         string                    `json:"name"`
	Description  string                    `json:"description"`
	Servings     float64                   `json:"servings"`
	Macros       Macros                    `json:"macros"`
	PrepTime     CookingTime               `json:"prepTime,omitempty"`
	TotalTime    CookingTime               `json:"totalTime,omitempty"`
	Ingredients  []RecipeIngredientSection `json:"ingredients"`
	Instructions []RecipeInstruction       `json:"instructions"`
	Notes        string                    `json:"notes,omitempty"`
	CreatedAt    time.Time                 `json:"-"`
	Version      int64                     `json:"version"`
}
