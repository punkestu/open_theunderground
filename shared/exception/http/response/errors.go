package response

import "github.com/punkestu/open_theunderground/shared/exception"

type FieldInvalids struct {
	Error []exception.Invalids `json:"exception"`
}

type ServerError struct {
	Error exception.Invalids `json:"exception"`
}

func NewServerError(message string) ServerError {
	return ServerError{
		Error: exception.New("internal", message),
	}
}
