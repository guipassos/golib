package covalent

import (
	"errors"
	"strings"

	"github.com/guipassos/golib/httphelper"
)

type (
	Covalent interface {
		GetWalletNFTs(in GetWalletNFTsIn) GetWalletNFTsOut
	}
	Options struct {
		APIURL     string
		APIKey     string
		HTTPHelper httphelper.Client
	}
	covalentImpl struct {
		apiURL     string
		apiKey     string
		httpHelper httphelper.Client
	}
)

func New(opt Options) (Covalent, error) {
	opt.SetDefault()
	if err := opt.Validate(); err != nil {
		return nil, err
	}
	return &covalentImpl{
		apiURL: opt.APIURL,
		apiKey: opt.APIKey,
	}, nil
}

func (o *Options) SetDefault() {
	if o.APIURL == "" {
		o.APIURL = DEFAULT_API_URL
	} else if !strings.HasSuffix(o.APIURL, "/") {
		o.APIURL = o.APIURL + "/"
	}
	if o.HTTPHelper == nil {
		o.HTTPHelper = httphelper.New(httphelper.Options{
			BaseURL: o.APIURL,
			Timeout: DEFAULT_TIMEOUT,
		})
		o.HTTPHelper.SetAuthBasicToHeader(o.APIKey, "")
	}
}

func (o *Options) Validate() error {
	if o.APIURL == "" {
		return errors.New("invalid api url")
	}
	if o.APIKey == "" {
		return errors.New("invalid api key")
	}
	return nil
}
