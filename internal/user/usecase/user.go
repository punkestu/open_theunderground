package usecase

import (
	"github.com/punkestu/open_theunderground/internal/user/entity/request"
	"github.com/punkestu/open_theunderground/internal/user/repo"
	"github.com/punkestu/open_theunderground/shared/domain"
	"github.com/punkestu/open_theunderground/shared/error/invalid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Repo *repo.User
}

func NewUser(repo *repo.User) *User {
	return &User{Repo: repo}
}

func (u *User) Login(req *request.Login) (user *domain.User, err error) {
	user, err = (*u.Repo).GetByUsername(req.Username)
	if user != nil && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		user, err = nil, invalid.New("password", "password is wrong")
	}
	return
}

func (u *User) Register(user *request.Register) (res *domain.User, err error) {
	res, err = (*u.Repo).Create(user.Fullname, user.Username, user.Password, user.Email)
	return
}

func (u *User) GetProfile(userId string) (user *domain.User, err error) {
	user, err = (*u.Repo).GetByID(userId)
	return
}
