package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.Handle("GET /api/v1/healthcheck", app.handleHealthcheck())

	router.Handle("GET /api/v1/measurements", app.handleShowAllMeasurements())
	router.Handle("POST /api/v1/measurements", app.handleCreateMeasurement())
	router.Handle("GET /api/v1/measurements/{id}", app.handleShowMeasurement())
	router.Handle("PUT /api/v1/measurements/{id}", app.handleUpdateMeasurement())

	router.Handle("POST /api/v1/foods", app.handleCreateFood())

	router.Handle("POST /api/v1/ingredients", app.handleCreateIngredient())

	router.Handle("POST /api/v1/recipes", app.handleCreateRecipe())
	router.HandleFunc("/api/v1/recipes", app.methodNotAllowedResponse)
	router.Handle("GET /api/v1/recipes/{id}", app.handleShowRecipe())
	router.HandleFunc("/api/v1/resipes/{id}", app.methodNotAllowedResponse)
	router.HandleFunc("/", app.notFoundResponse)

	return app.recoverPanic(router)
}
