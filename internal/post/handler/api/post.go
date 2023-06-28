package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/cerror/invalid"
	"github.com/punkestu/open_theunderground/internal/middleware/auth"
	"github.com/punkestu/open_theunderground/internal/post/entity/request"
	"github.com/punkestu/open_theunderground/internal/post/repo"
	"github.com/punkestu/open_theunderground/internal/post/usecase"
	"github.com/punkestu/open_theunderground/internal/user/entity/response"
	"net/http"
)

type Post struct {
	useCase *usecase.Post
}

func InitPost(r *fiber.App, repo repo.Post, middleware *auth.MidS) {
	useCase := Post{useCase: usecase.NewPost(&repo)}
	g := r.Group("/post")
	g.Use(middleware.IsAuth)
	g.Get("/", useCase.GetAll)
	g.Post("/", useCase.Create)
}

func (p *Post) GetAll(c *fiber.Ctx) error {
	mPosts, err := p.useCase.GetAll()
	if err != nil {
		if iErr := invalid.Parse(err); iErr != nil {
			return c.Status(http.StatusBadRequest).JSON(response.FieldInvalids{
				Error: []invalid.Invalids{
					*iErr,
				},
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	}
	return c.JSON(mPosts)
}

func (p *Post) Create(c *fiber.Ctx) error {
	var body request.Create
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	mPost, err := p.useCase.Create(body.Topic, c.Locals("userId").(string))
	if err != nil {
		if iErr := invalid.Parse(err); iErr != nil {
			return c.Status(http.StatusBadRequest).JSON(response.FieldInvalids{
				Error: []invalid.Invalids{
					*iErr,
				},
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(response.NewServerError(err.Error()))
	}
	return c.JSON(mPost)
}
