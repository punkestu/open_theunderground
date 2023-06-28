package domain

import "time"

type Post struct {
	ID        string       `json:"id"`
	Topic     string       `json:"topic"`
	Author    UserFiltered `json:"author"`
	CreatedAt time.Time    `json:"createdAt"`
}

type PostUseCase interface {
	GetAll() ([]*Post, error)
	GetByID(postId string) (*Post, error)
	GetByAuthor(authorId string) ([]*Post, error)
	Create(topic, authorId string) (*Post, error)
}
