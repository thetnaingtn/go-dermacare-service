package validate

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidPayload = errors.New("Invalid payload")
	ErrNotFound       = errors.New("Entity not found")
	ErrInvalidId      = errors.New("ID is not in a proper form")
)

type RequestError struct {
	Err    error
	Status int
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
			errorMsg := FieldError{Field: e.Field(), Error: msgForTag(e.Tag())}
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
