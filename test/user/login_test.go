package user_test

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/open_theunderground/internal/middleware/auth"
	mocks2 "github.com/punkestu/open_theunderground/internal/middleware/repo/mocks"
	"github.com/punkestu/open_theunderground/internal/user/entity/request"
	"github.com/punkestu/open_theunderground/internal/user/entity/response"
	"github.com/punkestu/open_theunderground/internal/user/handler/api"
	"github.com/punkestu/open_theunderground/internal/user/repo/mocks"
	"github.com/punkestu/open_theunderground/shared/domain"
	"github.com/punkestu/open_theunderground/shared/error/invalid"
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
	mock.On("GetByUsername", "minerv").Return(
		nil,
		invalid.New("username", "username is not found"),
	)
	mock.On("GetByUsername", "").Return(
		nil,
		errors.New("server error"),
	)
	const endpoint = "/user/login"
	jwtMock := *mocks2.NewJwtValidator(t)
	//IsValid(token string) (string, error)
	jwtMock.On("IsValid", "abcdefg").Return("user1234", nil)
	mids := auth.CreateMiddleware(&jwtMock)
	api.InitUser(app, &mock, mids)
	t.Run("Success", func(t *testing.T) {
		req, err := test.SendRequest(endpoint, request.Login{
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
		req, err := test.SendRequest(endpoint, request.Login{
			Username: "minerva",
			Password: "test123",
		}, nil)
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
		}, nil)
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
		}, nil)
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
