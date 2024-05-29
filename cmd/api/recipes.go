package main

import (
	"net/http"
)

func (app *application) handleCreateRecipe() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			app.writeJSON(w, http.StatusOK, envelope{"recipe": "new recipe"}, nil)
		},
	)
}

func (app *application) handleShowRecipe() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := app.writeJSON(w, http.StatusOK, envelope{"recipe": "show recipe"}, nil)
			if err != nil {
				app.serverErrorResponse(w, r, err)
			}
		},
	)
}
