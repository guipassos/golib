package httphelper

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
)

func (c *clientImpl) do(method string, request Request) (*http.Response, error) {
	bodyReader, err := GetBodyReader(request.Body)
	if err != nil {
		return nil, err
	}
	ctx := request.Context
	if ctx == nil {
		ctx = context.Background()
	}
	endpoint := strings.TrimPrefix(request.Endpoint, "/")
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+endpoint, bodyReader)
	if err != nil {
		return nil, err
	}
	for key, value := range c.header {
		req.Header.Set(key, value)
	}
	for key, value := range request.Header {
		req.Header.Set(key, value)
	}
	return c.httpClient.Do(req)
}

func (c *clientImpl) Get(request Request) *ResponseFail {
	response, err := c.do(http.MethodGet, request)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return DecodeResponse(response, request.Destination)
}

func (c *clientImpl) Post(request Request) *ResponseFail {
	response, err := c.do(http.MethodPost, request)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return DecodeResponse(response, request.Destination)
}

func (c *clientImpl) Patch(request Request) *ResponseFail {
	response, err := c.do(http.MethodPatch, request)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return DecodeResponse(response, request.Destination)
}

func (c *clientImpl) Put(request Request) *ResponseFail {
	response, err := c.do(http.MethodPut, request)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return DecodeResponse(response, request.Destination)
}

func (c *clientImpl) Delete(request Request) *ResponseFail {
	response, err := c.do(http.MethodDelete, request)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return DecodeResponse(response, request.Destination)
}

func (c clientImpl) SetAuthBasicToHeader(user, password string) {
	auth := []byte(user + ":" + password)
	c.header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString(auth)
}
