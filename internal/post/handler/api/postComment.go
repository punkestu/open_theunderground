package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/post/entity/request"
	"github.com/punkestu/open_theunderground/shared/exception"
	"net/http"
)

func (p *post) CreateComment(c *fiber.Ctx) error {
	var body request.PostComment
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(exception.NewInternal(err.Error()))
	}
	comment, err := p.useCase.CreateComment(body.PostID, c.Locals("userId").(string), body.Comment)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(exception.NewInternal(err.Error()))
	}
	return c.JSON(comment)
}
func (p *post) GetCommentByID(c *fiber.Ctx) error {
	commentID := c.Params("commentId")
	comment, err := p.useCase.GetCommentByID(commentID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(exception.NewInternal(err.Error()))
	}
	return c.JSON(comment)
}
func (p *post) GetCommentByPostID(c *fiber.Ctx) error {
	postID := c.Params("postId")
	comments, err := p.useCase.GetCommentByPostID(postID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(exception.NewInternal(err.Error()))
	}
	return c.JSON(comments)
}
