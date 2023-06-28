package usecase

import (
	"github.com/punkestu/open_theunderground/domain"
	"github.com/punkestu/open_theunderground/internal/post/repo"
)

type Post struct {
	Repo *repo.Post
}

func NewPost(repo *repo.Post) *Post {
	return &Post{Repo: repo}
}
func (p *Post) GetAll() (res *[]*domain.Post, err error) {
	res, err = (*p.Repo).GetAll()
	return
}
func (p *Post) GetByID(postId string) (res *domain.Post, err error) {
	res, err = (*p.Repo).GetByID(postId)
	return
}
func (p *Post) GetByAuthor(authorId string) (res *[]*domain.Post, err error) {
	res, err = (*p.Repo).GetByAuthor(authorId)
	return
}
func (p *Post) Create(topic, authorId string) (res *domain.Post, err error) {
	res, err = (*p.Repo).Create(topic, authorId)
	return
}
