package repo

import (
	"github.com/punkestu/open_theunderground/shared/domain"
)

type Post interface {
	GetAll() (*[]*domain.Post, error)
	GetByID(postId string) (*domain.Post, error)
	GetByAuthor(authorId string) (*[]*domain.Post, error)
	Create(topic string, authorId string) (*domain.Post, error)
	Update(topic string) (*domain.Post, error)
	GetLikeByPostID(postId string) ([]*domain.PostLike, error)
	GetLikeByPostAndUserID(postId, userId string) (*domain.PostLike, error)
	CreateLike(postId, userId string) (*domain.PostLike, error)
	DeleteLike(postId, userId string) error
}
