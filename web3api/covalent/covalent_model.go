package covalent

type (
	NFTData struct {
		TokenID      string `json:"token_id"`
		TokenBalance string `json:"token_balance"`
		TokenURL     string `json:"token_url"`
	}
	BalanceItem struct {
		ContractName         string    `json:"contract_name"`
		ContractTickerSymbol string    `json:"contract_ticker_symbol"`
		ContractAddress      string    `json:"contract_address"`
		LogoURL              string    `json:"logo_url"`
		Type                 string    `json:"type"`
		SupportsERC          []string  `json:"supports_erc"`
		NFTData              []NFTData `json:"nft_data"`
	}
	BalanceData struct {
		Items []BalanceItem `json:"items"`
	}
	GetBalanceOut struct {
		Data         *BalanceData `json:"data"`
		Error        bool         `json:"error"`
		ErrorMessage string       `json:"error_message"`
		ErrorCode    string       `json:"error_code"`
	}
)

type (
	GetExternalMetadataIn struct {
		ContractAddress string
		NFTID           string
		ChainID         int
	}
	GetExternalMetadataOut struct {
		Data         *ExternalData `json:"data"`
		Error        bool          `json:"error"`
		ErrorMessage string        `json:"error_message"`
		ErrorCode    string        `json:"error_code"`
	}
	ExternalData struct {
		Items []ExternalItem `json:"items"`
	}
	ExternalItem struct {
		NFTData []NFTData `json:"nft_data"`
	}
	ExternalNFTData struct {
		TokenID  string `json:"token_id"`
		TokenURL string `json:"token_url"`
	}
)


type (
	GetWalletNFTsIn struct {
		ChainID          int
		Wallet           string
		NFTAddressFilter string
	}
	GetWalletNFTsOut struct {
		List  *WalletNFTs `json:"list"`
		Error error       `json:"error"`
	}
	WalletNFTs []WalletNFT
	WalletNFT  struct {
		ContractAddress string      `json:"contract_address"`
		ContractName    string      `json:"contract_name"`
		ContractSymbol  string      `json:"contract_symbol"`
		ContractType    string      `json:"contract_type"`
		LogoURL         string      `json:"logo_url"`
		NFTData         NFTDataList `json:"nft_data"`
	}
	NFTDataList []NFTData
)

func (w WalletNFTs) IsEmpty() bool {
	return w == nil || len(w) == 0
}

func (w WalletNFTs) FilterByAddress(address string) WalletNFTs {
	if address == "" {
		return w
	}
	for _, nft := range w {
		if nft.ContractAddress == address {
			return WalletNFTs{
				nft,
			}
		}
	}
	return w
}
