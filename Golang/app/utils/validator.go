package utils

import (
	"github.com/go-playground/validator/v10"
)

type XValidator struct {
	validator *validator.Validate
}

var GlobalValidator = XValidator{validator: validate}

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.Error = true
			elem.FieldName = err.Field() // Export struct field name
			elem.Value = err.Value()     // Export field value

			switch err.Tag() {
			case "email":
				elem.Message = err.Field() + " must be a valid email address"
			case "min":
				elem.Message = err.Field() + " must be at least " + err.Param()
			case "max":
				elem.Message = err.Field() + " must be at most " + err.Param()
			default:
				elem.Message = err.Field() + " is required"
			}

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
