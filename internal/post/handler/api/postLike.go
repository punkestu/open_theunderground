package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/post/entity/request"
	"github.com/punkestu/open_theunderground/shared/exception"
	"net/http"
)

func (p *post) getLikeByPost(c *fiber.Ctx) error {
	postID := c.Params("postId")
	postLikes, err := (*p.useCase).GetLikeByPost(postID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(exception.NewInternal(err.Error()))
	}
	return c.JSON(postLikes)
}

func (p *post) toggleLike(c *fiber.Ctx) error {
	var body request.OnlyPostID
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(exception.NewInternal(err.Error()))
	}
	like, err := (*p.useCase).ToggleLike(body.PostID, c.Locals("userId").(string))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(exception.NewInternal(err.Error()))
	}
	if like != nil {
		return c.JSON(like)
	} else {
		return c.JSON(fiber.Map{
			"message": "removed",
		})
	}
}
