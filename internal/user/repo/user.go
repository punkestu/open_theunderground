package repo

import (
	"github.com/punkestu/open_theunderground/shared/domain"
)

type User interface {
	GetByID(userID string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	Create(fullname, username, password, email string) (*domain.User, error)
	Update(userId, fullname, username, email string) (*domain.User, error)
}
