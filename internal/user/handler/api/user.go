package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/middleware/auth"
	"github.com/punkestu/open_theunderground/internal/user/entity/request"
	"github.com/punkestu/open_theunderground/internal/user/entity/response"
	"github.com/punkestu/open_theunderground/internal/user/lib"
	"github.com/punkestu/open_theunderground/internal/user/repo"
	"github.com/punkestu/open_theunderground/internal/user/usecase"
	"github.com/punkestu/open_theunderground/shared/domain"
	"github.com/punkestu/open_theunderground/shared/error/invalid"
	"net/http"
)

type user struct {
	useCase *usecase.User
}

func InitUser(r *fiber.App, repo repo.User, midUser *auth.MidS) {
	useCase := user{usecase.NewUser(&repo)}
	g := r.Group("/user")
	g.Post("/login", useCase.login)
	g.Post("/register", useCase.register)
	g.Get("/me", midUser.IsAuth, useCase.profile)
}

func (u *user) login(c *fiber.Ctx) error {
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
	if token, err := lib.SignToken(*user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	} else {
		return c.JSON(response.JustToken{
			AuthToken: *token,
		})
	}
}

func (u *user) register(c *fiber.Ctx) error {
	var r request.Register
	if err := c.BodyParser(&r); err != nil {
		return err
	}
	mUser, err := u.useCase.Register(&r)
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
	if token, err := lib.SignToken(*mUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	} else {
		return c.JSON(response.JustToken{
			AuthToken: *token,
		})
	}
}

func (u *user) profile(c *fiber.Ctx) error {
	mUser, err := (*u.useCase.Repo).GetByID(c.Locals("userId").(string))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	}
	return c.JSON(domain.NewUserFiltered(mUser))
}
