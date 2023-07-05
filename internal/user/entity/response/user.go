package response

import (
	"github.com/punkestu/open_theunderground/shared/error/invalid"
)

type FieldInvalids struct {
	Error []invalid.Invalids `json:"error"`
}

type ServerError struct {
	Error invalid.Invalids `json:"error"`
}

func NewServerError(message string) ServerError {
	return ServerError{
		Error: invalid.New("internal", message),
	}
}
