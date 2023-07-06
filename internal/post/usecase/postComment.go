package usecase

import "github.com/punkestu/open_theunderground/shared/domain"

func (p *Post) CreateComment(postId, comment, userId string) (*domain.PostComment, error) {
	return (*p.Repo).CreateComment(postId, comment, userId)
}
func (p *Post) GetCommentByID(commentId string) (*domain.PostComment, error) {
	return (*p.Repo).GetCommentByID(commentId)
}
func (p *Post) GetCommentByPostID(postId string) ([]*domain.PostComment, error) {
	return (*p.Repo).GetCommentByPostID(postId)
}
