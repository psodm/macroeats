package data

import (
	"github.com/psodm/macroeats/internal/validator"
)

type Measurement struct {
	ID                      int64  `json:"id"`
	MeasurementName         string `json:"name"`
	MeasurementAbbreviation string `json:"abbreviation"`
}

func ValidateMeasurementUnit(v *validator.Validator, unit Measurement) {
	v.Check(unit.MeasurementName != "", "measurementName", "must be provided")
	v.Check(unit.MeasurementAbbreviation != "", "measurementAbbreviation", "must be provided")
}
