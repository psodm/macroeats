package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.Handle("GET /api/v1/healthcheck", app.handleHealthcheck())
	router.Handle("POST /api/v1/recipes", app.handleCreateRecipe())
	router.HandleFunc("/api/v1/recipes", app.methodNotAllowedResponse)
	router.Handle("GET /api/v1/recipes/{id}", app.handleShowRecipe())
	router.HandleFunc("/api/v1/resipes/{id}", app.methodNotAllowedResponse)
	router.HandleFunc("/", app.notFoundResponse)

	return app.recoverPanic(router)
}
