package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/post/entity/request"
	"github.com/punkestu/open_theunderground/shared/exception"
	excResp "github.com/punkestu/open_theunderground/shared/exception/http/response"
	"net/http"
)

func (p *post) getAll(c *fiber.Ctx) error {
	mPosts, err := p.useCase.GetAll()
	if err != nil {
		if iErr := exception.Parse(err); iErr != nil {
			return c.Status(http.StatusBadRequest).JSON(excResp.FieldInvalids{
				Error: []exception.Invalids{
					*iErr,
				},
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(excResp.NewServerError(err.Error()))
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
		if iErr := exception.Parse(err); iErr != nil {
			return c.Status(http.StatusBadRequest).JSON(excResp.FieldInvalids{
				Error: []exception.Invalids{
					*iErr,
				},
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(excResp.NewServerError(err.Error()))
	}
	return c.JSON(mPost)
}
