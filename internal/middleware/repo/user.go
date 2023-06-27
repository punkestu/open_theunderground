package repo

import "github.com/punkestu/open_theunderground/domain"

type User interface {
	GetByID(userID string) (*domain.User, error)
}
