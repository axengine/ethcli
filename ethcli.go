package ethcli

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmClient struct {
	*ethclient.Client
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

//func (cli *EvmClient) ChainID(ctx context.Context) (*big.Int, error) {
//	id, err := cli.ChainID(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return id, nil
//}
