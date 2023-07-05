package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	authResp "github.com/punkestu/open_theunderground/internal/middleware/entity/response"
	"github.com/punkestu/open_theunderground/internal/middleware/repo"
	"github.com/punkestu/open_theunderground/internal/user/entity/response"
	"github.com/punkestu/open_theunderground/shared/error/invalid"
	"net/http"
	"os"
)

type MidS struct {
	Repo *repo.User
}

func CreateMiddleware(Repo repo.User) *MidS {
	return &MidS{Repo: &Repo}
}

func (m *MidS) IsAuth(c *fiber.Ctx) error {
	if c.Get("Authorization") == "" {
		return c.Status(http.StatusUnauthorized).JSON(authResp.NewUnauthorized())
	}
	realToken := c.Get("Authorization")[7:]
	token, err := jwt.Parse(realToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if token == nil {
		return c.Status(http.StatusUnauthorized).JSON(authResp.NewUnauthorized())
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, _ := claims.GetSubject()
		_, err := (*m.Repo).GetByID(sub)
		if err != nil {
			if iErr := invalid.Parse(err); iErr != nil {
				return c.Status(http.StatusUnauthorized).JSON(authResp.NewInvalidToken())
			}
			return c.Status(http.StatusUnauthorized).JSON(response.NewServerError(err.Error()))
		}
		c.Locals("userId", sub)
		return c.Next()
	} else {
		return c.Status(http.StatusUnauthorized).JSON(authResp.NewInvalidToken())
	}
}
