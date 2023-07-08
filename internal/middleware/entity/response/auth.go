package response

import (
	"github.com/punkestu/open_theunderground/shared/exception"
)

type SingleError struct {
	Error exception.Invalids `json:"exception"`
}

func NewInvalidToken() SingleError {
	return SingleError{Error: exception.New("token", "invalid token")}
}

func NewUnauthorized() SingleError {
	return SingleError{Error: exception.New("token", "unauthorized")}
}
