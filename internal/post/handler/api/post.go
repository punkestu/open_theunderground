package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/post/entity/request"
	"github.com/punkestu/open_theunderground/internal/user/entity/response"
	"github.com/punkestu/open_theunderground/shared/error/invalid"
	"net/http"
)

func (p *post) getAll(c *fiber.Ctx) error {
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

func (p *post) create(c *fiber.Ctx) error {
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
