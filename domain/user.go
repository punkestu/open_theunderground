package domain

type User struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserUseCase interface {
	Login(username, password string) (*User, error)
	Register(fullname, username, password, email string) (*User, error)
	GetProfile(userID string) (*User, error)
}
