package validate

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidPayload    = errors.New("Invalid payload")
	ErrNotFound          = errors.New("Entity not found")
	ErrInvalidId         = errors.New("ID is not in a proper form")
	ErrFieldsValidation  = errors.New("Fields validation error")
	ErrIncorrectPassword = errors.New("Incorrect password")
)

type RequestError struct {
	Err    error
	Status int
}

type ErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

func NewRequestError(err error, status int) error {
	return &RequestError{
		Err:    err,
		Status: status,
	}
}

func (re *RequestError) Error() string {
	return re.Err.Error()
}

func IsRequestError(err error) bool {
	var re *RequestError
	return errors.As(err, &re)
}

type FieldError struct {
	Error string `json:"error"`
	Field string `json:"field"`
}

type FieldErrors []FieldError

func (fe FieldErrors) Error() string {
	fieldErrors, err := json.Marshal(fe)
	if err != nil {
		return err.Error()
	}

	return string(fieldErrors)
}

func IsFieldErrors(err error) bool {
	var fe FieldErrors
	return errors.As(err, &fe)
}

func GetFieldsValidationErrors(e error) FieldErrors {
	var ve validator.ValidationErrors
	if ok := errors.As(e, &ve); ok {
		errs := make([]FieldError, len(ve))
		for i, e := range ve {
			errorMsg := FieldError{Field: e.Field(), Error: msgForTag(e.Tag(), e.Field())}
			errs[i] = errorMsg
		}
		return errs
	}
	return nil
}

func (fe FieldErrors) Fields() map[string]string {
	fields := make(map[string]string)
	for _, f := range fe {
		fields[f.Field] = f.Error
	}
	return fields
}

func msgForTag(tag string, field string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "eqfield":
		return fmt.Sprintf("password and %s doesn't match", field)
	case "email":
		return fmt.Sprintf("%s has invalid format", field)
	default:
		return "Invalid tag"
	}
}
