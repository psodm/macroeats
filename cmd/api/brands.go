package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/psodm/macroeats/internal/data"
	"github.com/psodm/macroeats/internal/validator"
)

func (app *application) handleCreateBrand() http.Handler {
	type inputPayload struct {
		BrandName string `json:"brandName"`
	}
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			payload, err := decode[inputPayload](w, r)
			if err != nil {
				app.badRequestResponse(w, r, err)
				return
			}
			brand := data.Brand{
				BrandName: payload.BrandName,
			}
			v := validator.New()
			if data.ValidateBrand(v, brand); !v.Valid() {
				app.failedValidationResponse(w, r, v.Errors)
				return
			}
			err = app.models.Brands.Insert(&brand)
			if err != nil {
				switch {
				case errors.Is(err, data.ErrDuplicateRow):
					app.errorResponse(w, r, http.StatusConflict, err.Error())
					return
				default:
					app.serverErrorResponse(w, r, err)
					return
				}
			}

			headers := make(http.Header)
			headers.Set("Location", fmt.Sprintf("/api/v1/brands/%d", brand.ID))

			err = app.writeJSON(w, http.StatusCreated, envelope{"brand": brand}, headers)
			if err != nil {
				app.serverErrorResponse(w, r, err)
			}
		},
	)
}

func (app *application) handleShowBrand() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			name := app.readParam(r, "name")
			brand, err := app.models.Brands.GetByName(name)
			if err != nil {
				switch {
				case errors.Is(err, data.ErrRecordNotFound):
					app.notFoundResponse(w, r)
				default:
					app.logger.Error(err.Error())
					app.serverErrorResponse(w, r, err)
				}
				return
			}
			err = app.writeJSON(w, http.StatusOK, envelope{"brand": brand}, nil)
			if err != nil {
				app.logger.Error(err.Error())
				app.serverErrorResponse(w, r, err)
			}
		},
	)
}

func (app *application) handleShowAllBrands() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			brands, err := app.models.Brands.GetAll()
			if err != nil {
				switch {
				case errors.Is(err, data.ErrRecordNotFound):
					app.notFoundResponse(w, r)
				default:
					app.logger.Error(err.Error())
					app.serverErrorResponse(w, r, err)
				}
				return
			}
			err = app.writeJSON(w, http.StatusOK, envelope{"brands": brands}, nil)
			if err != nil {
				app.logger.Error(err.Error())
				app.serverErrorResponse(w, r, err)
			}
		},
	)
}
