package ethcli

import (
	"context"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmClient struct {
	*ethclient.Client

	_chainID atomic.Value
}

func New(rawurl string) (*EvmClient, error) {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return &EvmClient{
		Client: client,
	}, nil
}

func (cli *EvmClient) chainID() (*big.Int, error) {
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

// ID return chainId
func (cli *EvmClient) ID() *big.Int {
	id, err := cli.chainID()
	if err != nil {
		return big.NewInt(0)
	}
	return id
}
