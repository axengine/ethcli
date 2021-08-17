package ethcli

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"sync/atomic"
)

type ETHCli struct {
	*ethclient.Client

	_chainID atomic.Value
}

func New(rawurl string) *ETHCli {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Printf("New::Dial %v", err)
		return nil
	}
	return &ETHCli{
		Client: client,
	}
}

func (cli *ETHCli) chainID() *big.Int {
	chainID := cli._chainID.Load()
	if chainID != nil {
		return chainID.(*big.Int)
	}

	id, err := cli.ChainID(context.Background())
	if err != nil {
		log.Printf("chainID::ChainID %v", err)
		return nil
	}
	cli._chainID.Store(id)
	return id
}
