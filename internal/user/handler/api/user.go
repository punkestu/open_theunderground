package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/punkestu/open_theunderground/cerror/invalid"
	"github.com/punkestu/open_theunderground/domain"
	"github.com/punkestu/open_theunderground/internal/middleware/auth"
	"github.com/punkestu/open_theunderground/internal/user/entity/request"
	"github.com/punkestu/open_theunderground/internal/user/entity/response"
	"github.com/punkestu/open_theunderground/internal/user/repo"
	"github.com/punkestu/open_theunderground/internal/user/usecase"
	"net/http"
	"os"
	"time"
)

type User struct {
	useCase *usecase.User
}

func signToken(user domain.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(1 * time.Hour)},
		Subject:   user.ID,
	})
	if tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))); err != nil {
		return nil, err
	} else {
		return &tokenString, nil
	}
}

func InitUser(r *fiber.App, repo repo.User, midUser *auth.MidS) {
	useCase := User{usecase.NewUser(&repo)}
	g := r.Group("/user")
	g.Post("/login", useCase.login)
	g.Post("/register", useCase.register)
	g.Get("/me", midUser.IsAuth, useCase.profile)
}

func (u *User) login(c *fiber.Ctx) error {
	var r request.Login
	if err := c.BodyParser(&r); err != nil {
		return err
	}
	user, err := u.useCase.Login(&r)
	if err != nil {
		if iErr := invalid.Parse(err); iErr != nil {
			return c.Status(http.StatusUnauthorized).JSON(response.FieldInvalids{
				Error: []invalid.Invalids{
					*iErr,
				},
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	}
	if token, err := signToken(*user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	} else {
		return c.JSON(fiber.Map{
			"auth_token": *token,
		})
	}
}

func (u *User) register(c *fiber.Ctx) error {
	var r request.Register
	if err := c.BodyParser(&r); err != nil {
		return err
	}
	user, err := u.useCase.Register(&r)
	if err != nil {
		if iErr := invalid.Parse(err); iErr != nil {
			return c.Status(http.StatusUnauthorized).JSON(response.FieldInvalids{
				Error: []invalid.Invalids{
					*iErr,
				},
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	}
	if token, err := signToken(*user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	} else {
		return c.JSON(fiber.Map{
			"auth_token": *token,
		})
	}
}

func (u *User) profile(c *fiber.Ctx) error {
	mUser, err := (*u.useCase.Repo).GetByID(c.Locals("userId").(string))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	}
	return c.JSON(domain.NewUserFiltered(mUser))
}
