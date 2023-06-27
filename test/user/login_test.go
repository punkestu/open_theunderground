package user_test

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/cerror/invalid"
	"github.com/punkestu/open_theunderground/domain"
	"github.com/punkestu/open_theunderground/internal/user/entity/request"
	"github.com/punkestu/open_theunderground/internal/user/entity/response"
	"github.com/punkestu/open_theunderground/internal/user/handler/api"
	"github.com/punkestu/open_theunderground/internal/user/repo/mocks"
	"github.com/punkestu/open_theunderground/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	app := fiber.New()
	mock := *mocks.NewUser(t)
	mock.On("GetByUsername", "minerva").Return(&domain.User{
		ID:       "test1234",
		Fullname: "minerva the first",
		Username: "minerva",
		Password: "test1234",
		Email:    "minerva@mail.com",
	}, nil)
	mock.On("GetByUsername", "minerv").Return(
		nil,
		invalid.New("username", "username is not found"),
	)
	mock.On("GetByUsername", "").Return(
		nil,
		errors.New("server error"),
	)
	const endpoint = "/user/login"
	api.InitUser(app, &mock)
	t.Run("Success", func(t *testing.T) {
		req, err := test.SendRequest(endpoint, request.Login{
			Username: "minerva",
			Password: "test1234",
		})
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		resBody := domain.User{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, "test1234", resBody.ID)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
	t.Run("Password is wrong", func(t *testing.T) {
		req, err := test.SendRequest(endpoint, request.Login{
			Username: "minerva",
			Password: "test123",
		})
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		resBody := response.FieldInvalids{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		assert.Equal(t, "password is wrong", resBody.Error[0].Error())
	})
	t.Run("Username is wrong", func(t *testing.T) {
		req, err := test.SendRequest(endpoint, request.Login{
			Username: "minerv",
			Password: "test1234",
		})
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		resBody := response.FieldInvalids{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		assert.Equal(t, "username is not found", resBody.Error[0].Error())
	})
	t.Run("Server error", func(t *testing.T) {
		req, err := test.SendRequest(endpoint, request.Login{
			Username: "",
			Password: "test1234",
		})
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		resBody := response.ServerError{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.Equal(t, "server error", resBody.Error.Error())
	})
}
