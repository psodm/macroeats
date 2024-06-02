package main

import "net/http"

func (app *application) handleCreateIngredient() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("create ingredient"))
		},
	)
}
