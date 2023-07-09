package response

import "github.com/punkestu/open_theunderground/shared/domain"

type CreatePost struct {
	Post domain.Post `json:"post"`
}

type GetAll struct {
	Posts []*domain.Post `json:"posts"`
}
