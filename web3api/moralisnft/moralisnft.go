//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package moralisnft

import (
	"errors"
	"net/http"

	"github.com/rhizomplatform/golib/logger"
)

type (
	Options struct {
		ApiUrl     string
		ApiKey     string
		HttpClient *http.Client
	}
	MoralisNFT interface {
		GetList(addressWallet, addressNFT, chain, format, cursor string, limit int) ([]byte, error)
		GetWalletBalance(addressWallet, chain, tokenAddress string) ([]byte, error)
	}
	moralisNftImpl struct {
		apiUrl     string
		apiKey     string
		httpClient *http.Client
	}
)

func NewMoralisNFT(opt Options) MoralisNFT {
	opt.SetDefault()
	if err := opt.Validate(); err != nil {
		logger.Error(err)
		return nil
	}
	return &moralisNftImpl{
		apiUrl:     opt.ApiUrl,
		apiKey:     opt.ApiKey,
		httpClient: opt.HttpClient,
	}
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
