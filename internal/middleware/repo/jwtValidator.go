package repo

type JwtValidator interface {
	IsValid(token string) (string, error)
}
