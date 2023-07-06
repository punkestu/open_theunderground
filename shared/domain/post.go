package domain

import "time"

type Post struct {
	ID        string       `json:"id"`
	Topic     string       `json:"topic"`
	Author    UserFiltered `json:"author"`
	CreatedAt time.Time    `json:"createdAt"`
}

type PostLike struct {
	Post      *Post        `json:"post"`
	User      UserFiltered `json:"user"`
	CreatedAt time.Time    `json:"createdAt"`
}

type PostComment struct {
	ID        string       `json:"id"`
	User      UserFiltered `json:"user"`
	Comment   string       `json:"comment"`
	CreatedAt time.Time    `json:"createdAt"`
}

type PostUseCase interface {
	GetAll() ([]*Post, error)
	GetByID(postId string) (*Post, error)
	GetByAuthor(authorId string) ([]*Post, error)
	Create(topic, authorId string) (*Post, error)
	GetLikeByPost(postId string) ([]*PostLike, error)
	ToggleLike(postId, authorId string) (*PostLike, error)
	CreateComment(postId, comment, userId string) (*PostComment, error)
	GetCommentByID(commentId string) (*PostComment, error)
	GetCommentByPostID(postId string) ([]*PostComment, error)
}
