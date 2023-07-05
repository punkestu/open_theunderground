package repo

import (
	"github.com/punkestu/open_theunderground/shared/domain"
)

type User interface {
	GetByID(userID string) (*domain.User, error)
}
