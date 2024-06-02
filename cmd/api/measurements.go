package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/psodm/macroeats/internal/data"
	"github.com/psodm/macroeats/internal/validator"
)

type inputPayload struct {
	Name         string `json:"measurementName"`
	Abbreviation string `json:"measurementAbbreviation"`
}

func (app *application) handleCreateMeasurement() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			payload, err := decode[inputPayload](w, r)
			if err != nil {
				app.badRequestResponse(w, r, err)
				return
			}
			measurementUnit := data.Measurement{
				MeasurementName:         payload.Name,
				MeasurementAbbreviation: payload.Abbreviation,
			}

			v := validator.New()

			if data.ValidateMeasurementUnit(v, measurementUnit); !v.Valid() {
				app.failedValidationResponse(w, r, v.Errors)
				return
			}

			err = app.stores.measurementStore.Insert(r.Context(), &measurementUnit)
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
			headers.Set("Location", fmt.Sprintf("/api/v1/measurements/%d", measurementUnit.ID))

			err = app.writeJSON(w, http.StatusCreated, envelope{"measurement": measurementUnit}, headers)
			if err != nil {
				app.serverErrorResponse(w, r, err)
			}
		},
	)
}

func (app *application) handleShowMeasurement() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := app.readIDParam(r)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			measurement, err := app.stores.measurementStore.Get(r.Context(), id)
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
			err = app.writeJSON(w, http.StatusOK, envelope{"measurement": measurement}, nil)
			if err != nil {
				app.logger.Error(err.Error())
				app.serverErrorResponse(w, r, err)
			}
		},
	)
}

func (app *application) handleShowAllMeasurements() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			measurements, err := app.stores.measurementStore.GetAll(r.Context())
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
			err = app.writeJSON(w, http.StatusOK, envelope{"measurements": measurements}, nil)
			if err != nil {
				app.logger.Error(err.Error())
				app.serverErrorResponse(w, r, err)
			}
		},
	)
}

func (app *application) handleUpdateMeasurement() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := app.readIDParam(r)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			measurement, err := app.stores.measurementStore.Get(r.Context(), id)
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
			payload, err := decode[inputPayload](w, r)
			if err != nil {
				app.badRequestResponse(w, r, err)
				return
			}
			measurement.MeasurementName = payload.Name
			measurement.MeasurementAbbreviation = payload.Abbreviation
			v := validator.New()
			if data.ValidateMeasurementUnit(v, *measurement); !v.Valid() {
				app.failedValidationResponse(w, r, v.Errors)
				return
			}
			err = app.stores.measurementStore.Update(r.Context(), measurement)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}
			err = app.writeJSON(w, http.StatusOK, envelope{"measurement": measurement}, nil)
			if err != nil {
				app.logger.Error(err.Error())
				app.serverErrorResponse(w, r, err)
			}
		},
	)
}
