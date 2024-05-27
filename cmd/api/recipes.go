package main

import (
	"net/http"
	"time"

	"github.com/psodm/macroeats/internal/data"
	"github.com/psodm/macroeats/internal/validator"
)

func (app *application) handleCreateRecipe() http.Handler {
	type inputPayload struct {
		Name         string                         `json:"name"`
		Description  string                         `json:"description"`
		MealType     string                         `json:"mealType"`
		Cuisine      string                         `json:"cuisine"`
		Servings     float64                        `json:"servings"`
		Macros       data.Macros                    `json:"macros"`
		PrepTime     int64                          `json:"prepTime"`
		TotalTime    int64                          `json:"totalTime"`
		Ingredients  []data.RecipeIngredientSection `json:"ingredients"`
		Instructions []data.RecipeInstruction       `json:"instructions"`
		Notes        []string                       `json:"notes"`
	}
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			payload, err := decode[inputPayload](w, r)
			if err != nil {
				app.badRequestResponse(w, r, err)
				return
			}

			recipe := &data.Recipe{
				Name:         payload.Name,
				Description:  payload.Description,
				MealType:     payload.MealType,
				Cuisine:      payload.Cuisine,
				Servings:     payload.Servings,
				Macros:       payload.Macros,
				PrepTime:     data.CookingTime(payload.PrepTime),
				TotalTime:    data.CookingTime(payload.TotalTime),
				Ingredients:  payload.Ingredients,
				Instructions: payload.Instructions,
				Notes:        payload.Notes,
				CreatedAt:    time.Now(),
				Version:      1,
			}

			v := validator.New()

			if data.ValidateRecipe(v, recipe); !v.Valid() {
				app.failedValidationResponse(w, r, v.Errors)
				return
			}

			// fmt.Fprintf(w, "%+v\n", payload)
			app.writeJSON(w, http.StatusOK, envelope{"recipe": recipe}, nil)
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
				MealType:    "Dessert",
				Cuisine:     "Modern",
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
							{IngredientName: "Jelly Mix, Sugar Free", MeasurementQuantity: 36, MeasurementAbbreviation: "g"},
							{IngredientName: "Water", MeasurementQuantity: 500, MeasurementAbbreviation: "ml"},
							{IngredientName: "Greek Yoghurt, Fat Free", MeasurementQuantity: 500, MeasurementAbbreviation: "g"},
							{IngredientName: "Protein Powder, Unflavored", MeasurementQuantity: 60, MeasurementAbbreviation: "g"},
							{IngredientName: "Salt", MeasurementQuantity: 1, MeasurementAbbreviation: "pinch"},
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
				Notes: []string{
					"Suitable for meal prep. Keeps in the fridge 2-3 days. Can be made without salt, but a pinch of salt will enhance the sweetness",
				},
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
