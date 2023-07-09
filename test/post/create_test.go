package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/post/entity/request"
	"github.com/punkestu/open_theunderground/internal/post/entity/response"
	"github.com/punkestu/open_theunderground/internal/post/handler/api"
	"github.com/punkestu/open_theunderground/internal/post/repo/mocks"
	"github.com/punkestu/open_theunderground/shared/domain"
	"github.com/punkestu/open_theunderground/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	app := fiber.New()
	mock := *mocks.NewPost(t)
	//Create(topic string, authorId string) (*domain.Post, exception)
	mock.On("Create", "test", "user1234").Return(&domain.Post{
		ID:    "test1234",
		Topic: "test",
		Author: domain.UserFiltered{
			ID:       "user1234",
			Fullname: "user",
			Username: "user",
			Email:    "user",
		},
		CreatedAt: time.Now(),
	}, nil)
	const endpoint = "/post"
	api.InitPost(app, &mock, test.CreateAuthMock(t, "test_token", "user1234"))
	t.Run("Success", func(t *testing.T) {
		req, err := test.SendRequest(http.MethodPost, endpoint, request.Create{Topic: "test"}, map[string]string{
			"Authorization": "bearer test_token",
		})
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		resBody := response.CreatePost{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, "test1234", resBody.Post.ID)
	})
}
