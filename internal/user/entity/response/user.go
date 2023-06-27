package response

import (
	"github.com/punkestu/open_theunderground/cerror/invalid"
	"github.com/punkestu/open_theunderground/domain"
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

type UserFiltered struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserFiltered(user *domain.User) UserFiltered {
	return UserFiltered{
		ID:       user.ID,
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email,
	}
}
