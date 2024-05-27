package main

import (
	"fmt"
	"net/http"

	"github.com/psodm/macroeats/internal/data"
	"github.com/psodm/macroeats/internal/validator"
)

func (app *application) handleCreateFood() http.Handler {
	type inputPayload struct {
		Name               string  `json:"foodName"`
		BrandName          string  `json:"brandName"`
		ServingQuantity    float64 `json:"servingQuantity"`
		ServingMeasurement string  `json:"servingMeasurement"`
		Energy             float64 `json:"energy"`
		Calories           float64 `json:"calories"`
		Protein            float64 `json:"protein"`
		Carbohydrates      float64 `json:"carbohydrates"`
		Fat                float64 `json:"fat"`
	}
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			payload, err := decode[inputPayload](w, r)
			if err != nil {
				app.badRequestResponse(w, r, err)
				return
			}
			v := validator.New()
			measurement := data.Measurement{ID: 0, MeasurementName: "", MeasurementAbbreviation: payload.ServingMeasurement}
			macros := data.Macros{
				Energy:        payload.Energy,
				Calories:      payload.Calories,
				Protein:       payload.Protein,
				Carbohydrates: payload.Carbohydrates,
				Fat:           payload.Fat,
			}
			if data.ValidateMacros(v, macros); !v.Valid() {
				app.failedValidationResponse(w, r, v.Errors)
				return
			}
			macros.NormaliseMacrosEnergyAndCalories()
			food := data.FoodTx{
				ID:              0,
				FoodName:        payload.Name,
				BrandName:       payload.BrandName,
				ServingQuantity: payload.ServingQuantity,
				Measurement:     measurement,
				Macros:          macros,
			}
			err = app.models.Foods.InsertTx(&food)
			if err != nil {
				app.logger.Error(err.Error())
				app.serverErrorResponse(w, r, err)
				return
			}
			headers := make(http.Header)
			headers.Set("Location", fmt.Sprintf("/api/v1/foods/%d", food.ID))

			err = app.writeJSON(w, http.StatusCreated, envelope{"food": food}, headers)
			if err != nil {
				app.serverErrorResponse(w, r, err)
			}
		},
	)
}
