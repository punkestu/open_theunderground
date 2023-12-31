package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/user/entity/request"
	"github.com/punkestu/open_theunderground/internal/user/entity/response"
	"github.com/punkestu/open_theunderground/internal/user/handler/api"
	"github.com/punkestu/open_theunderground/internal/user/repo/mocks"
	"github.com/punkestu/open_theunderground/shared/domain"
	"github.com/punkestu/open_theunderground/shared/exception"
	excResp "github.com/punkestu/open_theunderground/shared/exception/http/response"
	"github.com/punkestu/open_theunderground/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRegister(t *testing.T) {
	app := fiber.New()
	mock := *mocks.NewUser(t)
	mock.On("Create", "the minerva", "username", "test1234", "minerva@mail.com").Return(&domain.User{
		ID:       "test1234",
		Fullname: "the minerva",
		Username: "minerva",
		Password: "test1234",
		Email:    "minerva@mail.com",
	}, nil)
	mock.On("Create", "the minerva", "minerva", "test1234", "minerva@mail.com").Return(nil, exception.New("username", "username is used"))
	const endpoint = "/user/register"
	api.InitUser(app, &mock, test.CreateAuthMock(t, "test_token", "user1234"))
	t.Run("Success", func(t *testing.T) {
		req, err := test.SendRequest(http.MethodPost, endpoint, request.Register{
			Fullname: "the minerva",
			Username: "username",
			Password: "test1234",
			Email:    "minerva@mail.com",
		}, nil)
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		resBody := response.JustToken{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.NotNil(t, resBody.AuthToken)
	})
	t.Run("Username is used", func(t *testing.T) {
		req, err := test.SendRequest(http.MethodPost, endpoint, request.Register{
			Fullname: "the minerva",
			Username: "minerva",
			Password: "test1234",
			Email:    "minerva@mail.com",
		}, nil)
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		resBody := excResp.FieldInvalids{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, "username is used", resBody.Error[0].Error())
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
