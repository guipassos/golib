package moralisnft

import (
	"errors"
	"net/http"
)

type moralisNftImpl struct {
	apiUrl     string
	apiKey     string
	httpClient *http.Client
}

type Options struct {
	ApiUrl     string
	ApiKey     string
	HttpClient *http.Client
}

func (o *Options) Validate() error {
	if o.ApiUrl == "" {
		return errors.New("invalid api url")
	}
	if o.ApiKey == "" {
		return errors.New("invalid api key")
	}
	if o.HttpClient == nil {
		return errors.New("invalid http client")
	}
	return nil
}

func (o *Options) SetDefault() {
	if o.ApiUrl == "" {
		o.ApiUrl = "https://deep-index.moralis.io/api/v2/"
	} else {
		if o.ApiUrl[len(o.ApiUrl)-1:] != "/" {
			o.ApiUrl = o.ApiUrl + "/"
		}
	}
	if o.HttpClient == nil {
		o.HttpClient = NewHttpClient(o.ApiKey)
	}
}
