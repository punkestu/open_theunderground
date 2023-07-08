package jwtValidator

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	authResp "github.com/punkestu/open_theunderground/internal/middleware/entity/response"
	"github.com/punkestu/open_theunderground/internal/middleware/repo"
	"github.com/punkestu/open_theunderground/internal/user/entity/response"
	"github.com/punkestu/open_theunderground/shared/exception"
	"os"
)

type Validator struct {
	userRepo *repo.User
}

func NewValidator(userRepo repo.User) *Validator {
	return &Validator{
		userRepo: &userRepo,
	}
}

func (v *Validator) IsValid(realToken string) (string, error) {
	token, err := jwt.Parse(realToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if token == nil {
		return "", authResp.NewUnauthorized().Error
	}
	if err != nil {
		return "", response.NewServerError(err.Error()).Error
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, _ := claims.GetSubject()
		_, err := (*v.userRepo).GetByID(sub)
		if err != nil {
			if iErr := exception.Parse(err); iErr != nil {
				return "", authResp.NewInvalidToken().Error
			}
			return "", response.NewServerError(err.Error()).Error
		}
		return sub, nil
	}
	return "", authResp.NewInvalidToken().Error
}
