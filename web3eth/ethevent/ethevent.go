//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package ethevent

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rhizomplatform/golib/web3eth/ethclient"
)

type (
	Options struct {
		Client ethclient.Client
	}
	Event interface {
		SubscribingLogsBlocks(from *big.Int) error
	}
	implEvent struct {
		client ethclient.Client
	}
)

func NewEvent(opts Options) Event {
	return &implEvent{
		client: opts.Client,
	}
}

func (e *implEvent) SubscribingLogsBlocks(from *big.Int) error {
	ctx := context.Background()
	query := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   from,
	}
	ethLogs := make(chan types.Log)
	sub, err := e.client.SubscribeFilterLogs(ctx, query, ethLogs)
	if err != nil {
		return err
	}
	block, err := e.client.BlockByNumber(ctx, from)
	body := block.Body()
	txs := body.Transactions
	txs[0].Type()
	for {
		select {
		case err := <-sub.Err():
			return err
		case ethLog := <-ethLogs:
			fmt.Println(ethLog)
		}
	}
}
