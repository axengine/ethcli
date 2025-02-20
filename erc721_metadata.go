package ethcli

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	openzeppelinERC721MetadataAbi = `[{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
)

func (cli *EvmClient) ERC721Name(ctx context.Context, token string, blockNumber *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721MetadataAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("name")

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return "", err
	}

	results, err := ins.Unpack("name", bz)
	if err != nil {
		return "", err
	}

	return results[0].(string), nil
}

func (cli *EvmClient) ERC721Symbol(ctx context.Context, token string, blockNumber *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721MetadataAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("symbol")

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return "", err
	}

	results, err := ins.Unpack("symbol", bz)
	if err != nil {
		return "", err
	}

	return results[0].(string), nil
}

// ERC721TokenURI
// for ERC721Metadata && ERC721URIStorage
//func (cli *EvmClient) ERC721TokenURI(token string, tokenId *big.Int) (string, error)
