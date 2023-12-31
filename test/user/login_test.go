package user_test

import (
	"errors"
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
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	app := fiber.New()
	mock := *mocks.NewUser(t)
	password, err := bcrypt.GenerateFromPassword([]byte("test1234"), 10)
	if err != nil {
		log.Fatal("system is broke, it cannot encrypt")
	}
	mock.On("GetByUsername", "minerva").Return(&domain.User{
		ID:       "test1234",
		Fullname: "minerva the first",
		Username: "minerva",
		Password: string(password),
		Email:    "minerva@mail.com",
	}, nil)
	mock.On("GetByUsername", "wrong_username").Return(
		nil,
		exception.New("username", "username is not found"),
	)
	mock.On("GetByUsername", "").Return(
		nil,
		errors.New("server exception"),
	)
	const endpoint = "/user/login"
	api.InitUser(app, &mock, test.CreateAuthMock(t, "test_token", "user1234"))
	t.Run("Success", func(t *testing.T) {
		req, err := test.SendRequest(http.MethodPost, endpoint, request.Login{
			Username: "minerva",
			Password: "test1234",
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
	t.Run("Password is wrong", func(t *testing.T) {
		req, err := test.SendRequest(http.MethodPost, endpoint, request.Login{
			Username: "minerva",
			Password: "test123",
		}, nil)
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		resBody := excResp.FieldInvalids{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		assert.Equal(t, "password is wrong", resBody.Error[0].Error())
	})
	t.Run("Username is wrong", func(t *testing.T) {
		req, err := test.SendRequest(http.MethodPost, endpoint, request.Login{
			Username: "wrong_username",
			Password: "test1234",
		}, nil)
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		resBody := excResp.FieldInvalids{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		assert.Equal(t, "username is not found", resBody.Error[0].Error())
	})
	t.Run("Server exception", func(t *testing.T) {
		req, err := test.SendRequest(http.MethodPost, endpoint, request.Login{
			Username: "",
			Password: "test1234",
		}, nil)
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		resBody := excResp.ServerError{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.Equal(t, "server exception", resBody.Error.Error())
	})
}
