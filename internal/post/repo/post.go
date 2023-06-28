package repo

import "github.com/punkestu/open_theunderground/domain"

type Post interface {
	GetAll() (*[]*domain.Post, error)
	GetByID(postId string) (*domain.Post, error)
	GetByAuthor(authorId string) (*[]*domain.Post, error)
	Create(topic string, authorId string) (*domain.Post, error)
	Update(topic string) (*domain.Post, error)
}
