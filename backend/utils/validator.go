package utils

import "github.com/go-playground/validator/v10"

func NewValidator() *validator.Validate {
	// Create a single instance of Validate
	validate := validator.New()

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}
