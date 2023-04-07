//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package moralisnft

import "github.com/rhizomplatform/golib/logger"

type MoralisNFT interface {
	GetList(addressWallet, addressNFT, chain, format, cursor string, limit int) ([]byte, error)
	GetWalletBalance(addressWallet, chain, tokenAddress string) ([]byte, error)
}

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
