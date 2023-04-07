package moralisnft

import (
	"net/http"
	"time"
)

type HttpTransport struct {
	ApiKey       string
	roundTripper http.RoundTripper
}

func NewHttpTransport(apiKey string) *HttpTransport {
	return &HttpTransport{
		ApiKey: apiKey,
	}
}

func (t *HttpTransport) Transport() http.RoundTripper {
	if t.roundTripper != nil {
		return t.roundTripper
	}
	return http.DefaultTransport
}

func (t *HttpTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("X-API-Key", t.ApiKey)
	return t.Transport().RoundTrip(r)
}

func NewHttpClient(apiKey string) *http.Client {
	return &http.Client{
		Timeout:   5 * time.Second,
		Transport: NewHttpTransport(apiKey),
	}
}
