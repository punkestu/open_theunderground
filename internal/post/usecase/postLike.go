package usecase

import (
	"github.com/punkestu/open_theunderground/shared/domain"
)

func (p *Post) GetLikeByPost(postId string) ([]*domain.PostLike, error) {
	return (*p.Repo).GetLikeByPostID(postId)
}

func (p *Post) ToggleLike(postId, userId string) (*domain.PostLike, error) {
	like, err := (*p.Repo).GetLikeByPostAndUserID(postId, userId)
	if err != nil {
		return nil, err
	}
	if like != nil {
		err := (*p.Repo).DeleteLike(postId, userId)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		newLike, err := (*p.Repo).CreateLike(postId, userId)
		if err != nil {
			return nil, err
		}
		return newLike, nil
	}
}
