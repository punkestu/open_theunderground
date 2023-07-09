package test

import (
	"bytes"
	"encoding/json"
	"github.com/punkestu/open_theunderground/internal/middleware/auth"
	mocks2 "github.com/punkestu/open_theunderground/internal/middleware/repo/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GetBody(resp *http.Response, dst any) error {
	resBodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
		return err
	}
	err = json.Unmarshal(resBodyByte, &dst)
	if err != nil {
		println(err.Error())
	}
	return nil
}

func SendRequest(method, endpoint string, body any, headers map[string]string) (req *http.Request, err error) {
	reqBody, err := json.Marshal(&body)
	req = httptest.NewRequest(method, endpoint, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	for name, header := range headers {
		req.Header.Set(name, header)
	}
	return
}

func CreateAuthMock(t *testing.T, forToken, resId string) *auth.MidS {
	jwtMock := *mocks2.NewJwtValidator(t)
	//IsValid(token string) (string, exception)
	jwtMock.On("IsValid", forToken).Return(resId, nil)
	return auth.CreateMiddleware(&jwtMock)
}
