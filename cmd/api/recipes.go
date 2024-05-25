package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/psodm/macroeats/internal/data"
)

func (app *application) handleCreateRecipe() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "create a new recipe")
		},
	)
}

func (app *application) handleShowRecipe() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := app.readIDParam(r)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			recipe := data.Recipe{
				ID:          id,
				Name:        "High Protein Mousse Bowl",
				Description: "High protein, low fat and low carb mousse bowl",
				Servings:    3,
				Macros: data.Macros{
					Energy:        715.4,
					Calories:      171,
					Protein:       29,
					Carbohydrates: 12.4,
					Fat:           0.9,
				},
				PrepTime:  5,
				TotalTime: 185,
				Ingredients: []data.RecipeIngredientSection{
					{SectionName: "Mousse Mixture",
						Ingredients: []data.RecipeIngredient{
							{IngredientName: "Jelly Mix, Sugar Free", MeasurementQuantity: 36, MeasurementDescription: "g"},
							{IngredientName: "Water", MeasurementQuantity: 500, MeasurementDescription: "ml"},
							{IngredientName: "Greek Yoghurt, Fat Free", MeasurementQuantity: 500, MeasurementDescription: "g"},
							{IngredientName: "Protein Powder, Unflavored", MeasurementQuantity: 60, MeasurementDescription: "g"},
							{IngredientName: "Salt", MeasurementQuantity: 1, MeasurementDescription: "pinch"},
						},
					},
				},
				Instructions: []data.RecipeInstruction{
					{Step: 1, Description: "Dissolve jelly mix in boiling water and set aside"},
					{Step: 2, Description: "Combine greek yoghurt, protein powder and salt in a blender"},
					{Step: 3, Description: "Add the jelly solution to the blender"},
					{Step: 4, Description: "Blend on low until the mixture is incorporated and smooth"},
					{Step: 5, Description: "Divide the mixture between 3 bowls"},
					{Step: 6, Description: "Cover each bowl and chill in the fridge for at least 3 hours"},
				},
				Notes:     "Suitable for meal prep. Keeps in the fridge 2-3 days. Can be made without salt, but a pinch of salt will enhance the sweetness",
				CreatedAt: time.Now(),
				Version:   1,
			}
			err = app.writeJSON(w, http.StatusOK, envelope{"recipe": recipe}, nil)
			if err != nil {
				app.serverErrorResponse(w, r, err)
			}
		},
	)
}
