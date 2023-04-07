//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package httphelper

import (
	"context"
	"net/http"
	"strings"
	"time"
)

type (
	Options struct {
		BaseURL    string
		HttpClient *http.Client
		Timeout    time.Duration
		Header     map[string]string
	}
	Client interface {
		Get(request Request) *ResponseFail
		Post(request Request) *ResponseFail
		Patch(request Request) *ResponseFail
		Put(request Request) *ResponseFail
		Delete(request Request) *ResponseFail
		SetAuthBasicToHeader(user, password string)
	}
	Request struct {
		Context     context.Context
		Endpoint    string
		Header      map[string]string
		Body        interface{}
		Destination interface{}
	}
	ResponseFail struct {
		Err        error
		StatusCode int
		Body       []byte
		Data       map[string]interface{}
		Header     http.Header
	}
	clientImpl struct {
		baseURL    string
		header     map[string]string
		httpClient *http.Client
	}
)

func New(opts Options) Client {
	if opts.BaseURL != "" && !strings.HasSuffix(opts.BaseURL, "/") {
		opts.BaseURL += "/"
	}
	if opts.Header == nil {
		opts.Header = map[string]string{
			"Content-Type": "application/json",
		}
	}
	httpClient := opts.HttpClient
	if opts.HttpClient == nil {
		httpClient = http.DefaultClient
	}
	if opts.Timeout > 0 {
		httpClient.Timeout = opts.Timeout
	}
	return &clientImpl{
		httpClient: httpClient,
		baseURL:    opts.BaseURL,
		header:     opts.Header,
	}
}
