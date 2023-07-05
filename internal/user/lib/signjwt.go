package lib

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/punkestu/open_theunderground/shared/domain"
	"os"
	"time"
)

func SignToken(mUser domain.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(1 * time.Hour)},
		Subject:   mUser.ID,
	})
	if tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))); err != nil {
		return nil, err
	} else {
		return &tokenString, nil
	}
}
