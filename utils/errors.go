package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/varomnrg/money-tracker/model"
)

func JSONError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	var errorResponse model.ErrorResponse
	errorResponse.Error = err.Error()
	json.NewEncoder(w).Encode(errorResponse)
}

func JSONErrorMap(w http.ResponseWriter, err map[string]string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err,
	})
}

func MapErrors(err error, validate *validator.Validate) map[string]string {
	errs := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		message := err.Error()

		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", err.Field())
		case "email":
			message = fmt.Sprintf("%s is not valid email", err.Field())
		case "min":
			message = fmt.Sprintf("%s must be at least %s character", err.Field(), err.Param())
		case "max":
			message = fmt.Sprintf("%s must be at most %s character", err.Field(), err.Param())
		case "alphanum":
			message = fmt.Sprintf("%s must be alphanumeric", err.Field())
		default:
			message = fmt.Sprintf("%s is not valid", err.Field())
		}

		errs[err.Field()] = message
	}

	return errs
}
