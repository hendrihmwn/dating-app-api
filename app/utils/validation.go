package utils

import "github.com/go-playground/validator/v10"

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "max":
		return "Max character is " + fe.Param()
	case "min":
		return "Min character is " + fe.Param()
	}
	return "Unknown error"
}
