package user

import (
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

func TestRegister(t *testing.T) {
	app := fiber.New()
	mock := *mocks.NewUser(t)
	mock.On("Create", "the minerva", "dobberman", "test1234", "minerva@mail.com").Return(&domain.User{
		ID:       "test1234",
		Fullname: "the minerva",
		Username: "minerva",
		Password: "test1234",
		Email:    "minerva@mail.com",
	}, nil)
	mock.On("Create", "the minerva", "minerva", "test1234", "minerva@mail.com").Return(nil, invalid.New("username", "username is used"))
	const endpoint = "/user/register"
	api.InitUser(app, &mock)
	t.Run("Success", func(t *testing.T) {
		req, err := test.SendRequest(endpoint, request.Register{
			Fullname: "the minerva",
			Username: "dobberman",
			Password: "test1234",
			Email:    "minerva@mail.com",
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
	t.Run("Username is used", func(t *testing.T) {
		req, err := test.SendRequest(endpoint, request.Register{
			Fullname: "the minerva",
			Username: "minerva",
			Password: "test1234",
			Email:    "minerva@mail.com",
		})
		assert.Nil(t, err)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		resBody := response.FieldInvalids{}
		err = test.GetBody(resp, &resBody)
		assert.Nil(t, err)
		assert.Equal(t, "username is used", resBody.Error[0].Error())
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
