package server

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "email":
		return "Should be email"
	}

	return "Unknown error"
}

func TransformErrorMessage(err error) []ErrorMessage {
	var validatorErrs validator.ValidationErrors

	if errors.As(err, &validatorErrs) {
		out := make([]ErrorMessage, len(validatorErrs))

		for i, fe := range validatorErrs {
			out[i] = ErrorMessage{fe.Field(), getErrorMessage(fe)}
		}

		return out
	}

	return nil
}