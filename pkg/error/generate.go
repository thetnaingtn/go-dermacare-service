package error

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Msg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetFieldsValidationErrors(e error) []Msg {
	var ve validator.ValidationErrors
	if ok := errors.As(e, &ve); ok {
		errs := make([]Msg, len(ve))
		for i, e := range ve {
			errorMsg := Msg{Field: e.Field(), Message: msgForTag(e.Tag())}
			errs[i] = errorMsg
		}
		return errs
	}
	return nil
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	default:
		return "Invalid tag"
	}
}
