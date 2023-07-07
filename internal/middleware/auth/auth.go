package auth

import (
	"github.com/gofiber/fiber/v2"
	authResp "github.com/punkestu/open_theunderground/internal/middleware/entity/response"
	"github.com/punkestu/open_theunderground/internal/middleware/repo"
	"net/http"
)

type MidS struct {
	validator *repo.JwtValidator
}

func CreateMiddleware(validator repo.JwtValidator) *MidS {
	return &MidS{validator: &validator}
}

func (m *MidS) IsAuth(c *fiber.Ctx) error {
	if c.Get("Authorization") == "" {
		return c.Status(http.StatusUnauthorized).JSON(authResp.NewUnauthorized())
	}
	realToken := c.Get("Authorization")[7:]
	result, err := (*m.validator).IsValid(realToken)
	if err == nil {
		c.Locals("userId", result)
		return c.Next()
	} else {
		return c.Status(http.StatusUnauthorized).JSON(err)
	}
}
