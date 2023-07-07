package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/middleware/auth"
	mocks2 "github.com/punkestu/open_theunderground/internal/middleware/repo/mocks"
	"github.com/punkestu/open_theunderground/internal/post/entity/request"
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
	//Create(topic string, authorId string) (*domain.Post, error)
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
	jwtMock := *mocks2.NewJwtValidator(t)
	//IsValid(token string) (string, error)
	jwtMock.On("IsValid", "abcdefg").Return("user1234", nil)
	const endpoint = "/post"
	mids := auth.CreateMiddleware(&jwtMock)
	api.InitPost(app, &mock, mids)
	t.Run("Success", func(t *testing.T) {
		req, err := test.SendRequest(endpoint, request.Create{Topic: "test"}, map[string]string{
			"Authorization": "bearer abcdefg",
		})
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		resBody := domain.Post{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, "test1234", resBody.ID)
	})
}