package covalent

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/rhizomplatform/golib/httphelper"
	"github.com/rhizomplatform/golib/logger"
)

const (
	BALANCE_ITEM_NFT_TYPE = "nft"
	ERC721_TYPE           = "erc721"
	ERC1155_TYPE          = "erc1155"
)

func (c covalentImpl) GetWalletNFTs(in GetWalletNFTsIn) (out GetWalletNFTsOut) {
	balance, err := c.getBalanceFromHttp(in)
	if err != nil {
		return GetWalletNFTsOut{
			List:  nil,
			Error: err,
		}
	}
	nftList := c.extractNFTListFromBalance(balance)
	nftList = nftList.FilterByAddress(in.NFTAddressFilter)
	nftList = c.getNFTListMetadata(in.ChainID, nftList)
	return GetWalletNFTsOut{
		List:  &nftList,
		Error: nil,
	}
}

func (c covalentImpl) getNFTListMetadata(chainID int, nftList WalletNFTs) WalletNFTs {
	var wg sync.WaitGroup
	for k := 0; k < len(nftList); k++ {
		nft := &nftList[k]
		for j := 0; j < len(nft.NFTData); j++ {
			nftData := &nft.NFTData[j]
			wg.Add(1)
			go func(wg *sync.WaitGroup, nftData *NFTData, c covalentImpl) {
				defer wg.Done()
				request := GetExternalMetadataIn{
					ContractAddress: nft.ContractAddress,
					ChainID:         chainID,
					NFTID:           nftData.TokenID,
				}
				httpResponse, err := c.getExternalMetadataFromHttp(request)
				if err != nil {
					logger.Error("failed to get external, request: ", request, " metadata: ", err)
				}
				validResponse := httpResponse != nil && httpResponse.Data != nil && httpResponse.Error == false && len(httpResponse.Data.Items) > 0
				if validResponse {
					metadata := httpResponse.Data.Items[0].NFTData
					if len(metadata) > 0 {
						nftData.TokenURL = metadata[0].TokenURL
					}
				}
			}(&wg, nftData, c)
		}
	}
	wg.Wait()
	return nftList
}

func (c covalentImpl) getBalanceFromHttp(in GetWalletNFTsIn) (*GetBalanceOut, error) {
	var (
		balance = &GetBalanceOut{}
		fail    = c.httpHelper.Get(httphelper.Request{
			Context:     context.Background(),
			Endpoint:    fmt.Sprintf("%d/address/%s/balances_v2/?nft=true&no-nft-fetch=true", in.ChainID, in.Wallet),
			Destination: balance,
		})
	)
	if fail != nil && fail.Err != nil {
		return nil, fail.Err
	}
	return balance, nil
}

func (c covalentImpl) getExternalMetadataFromHttp(in GetExternalMetadataIn) (*GetExternalMetadataOut, error) {
	var (
		nftMetadata = &GetExternalMetadataOut{}
		fail        = c.httpHelper.Get(httphelper.Request{
			Context:     context.Background(),
			Endpoint:    fmt.Sprintf("%d/tokens/%s/nft_metadata/%s/", in.ChainID, in.ContractAddress, in.NFTID),
			Destination: nftMetadata,
		})
	)
	if fail != nil && fail.Err != nil {
		return nil, fail.Err
	}
	return nftMetadata, nil
}

func (c covalentImpl) extractNFTListFromBalance(balance *GetBalanceOut) WalletNFTs {
	if balance == nil || balance.Data == nil {
		return WalletNFTs{}
	}
	var (
		conType string
		nftData []NFTData
		nftItem WalletNFT
		nftList = make([]WalletNFT, 0, len(balance.Data.Items))
	)
	for _, item := range balance.Data.Items {
		if item.Type == BALANCE_ITEM_NFT_TYPE {
			conType = CONTRACT_OTHER_TYPE
			for _, ercType := range item.SupportsERC {
				switch strings.ToLower(ercType) {
				case ERC721_TYPE:
					conType = CONTRACT_ERC721_TYPE
					break
				case ERC1155_TYPE:
					conType = CONTRACT_ERC1155_TYPE
					break
				}
			}
			nftData = make([]NFTData, len(item.NFTData))
			for k := 0; k < len(item.NFTData); k++ {
				nftData[k] = NFTData{
					TokenID:      item.NFTData[k].TokenID,
					TokenBalance: item.NFTData[k].TokenBalance,
					TokenURL:     item.NFTData[k].TokenURL,
				}
			}
			nftItem = WalletNFT{
				ContractAddress: item.ContractAddress,
				ContractName:    item.ContractName,
				ContractSymbol:  item.ContractTickerSymbol,
				ContractType:    conType,
				LogoURL:         item.LogoURL,
				NFTData:         nftData,
			}
			nftList = append(nftList, nftItem)
		}
	}
	return nftList
}
