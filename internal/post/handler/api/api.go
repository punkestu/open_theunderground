package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/middleware/auth"
	"github.com/punkestu/open_theunderground/internal/post/repo"
	"github.com/punkestu/open_theunderground/internal/post/usecase"
)

type post struct {
	useCase *usecase.Post
}

func InitPost(r *fiber.App, repo repo.Post, middleware *auth.MidS) {
	useCase := post{useCase: usecase.NewPost(&repo)}
	g := r.Group("/post")
	g.Use(middleware.IsAuth)
	g.Get("/", useCase.getAll)
	g.Post("/", useCase.create)
	g.Post("/like", useCase.toggleLike)
	g.Get("/like/:postId", useCase.getLikeByPost)
	g.Post("/comment", useCase.CreateComment)
	g.Get("/comment/:commentId", useCase.GetCommentByID)
	g.Get("/:postId/comment", useCase.GetCommentByPostID)
}
