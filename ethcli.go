package ethcli

import (
	"context"
	"github.com/axengine/ethcli/eth/ethclient"
	"math/big"
	"sync/atomic"
)

type ETHCli struct {
	*ethclient.Client

	_chainID atomic.Value
}

func New(rawurl string) (*ETHCli, error) {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return &ETHCli{
		Client: client,
	}, nil
}

func (cli *ETHCli) chainID() (*big.Int, error) {
	chainID := cli._chainID.Load()
	if chainID != nil {
		return chainID.(*big.Int), nil
	}

	id, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	cli._chainID.Store(id)
	return id, nil
}
