package ethcli

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	openzeppelinERC721URIStorageAbi = `[{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
)

// ERC721TokenURI
// for ERC721Metadata && ERC721URIStorage
func ERC721TokenURI(ctx context.Context, cli *ethclient.Client, token string, tokenId *big.Int, blockNumber *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721URIStorageAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("tokenURI", tokenId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return "", err
	}

	results, err := ins.Unpack("tokenURI", bz)
	if err != nil {
		return "", err
	}

	return results[0].(string), nil
}
