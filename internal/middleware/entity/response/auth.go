package response

import (
	"github.com/punkestu/open_theunderground/shared/error/invalid"
)

type SingleError struct {
	Error invalid.Invalids `json:"error"`
}

func NewInvalidToken() SingleError {
	return SingleError{Error: invalid.New("token", "invalid token")}
}

func NewUnauthorized() SingleError {
	return SingleError{Error: invalid.New("token", "unauthorized")}
}
