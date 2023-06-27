package request

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
