package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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

func SendRequest(endpoint string, body any) (req *http.Request, err error) {
	reqBody, err := json.Marshal(&body)
	req = httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	return
}
