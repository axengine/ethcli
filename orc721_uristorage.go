package ethcli

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

var (
	openzeppelinERC721URIStorageAbi = `[{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
)

// ORC721TokenURI
// for ERC721Metadata && ERC721URIStorage
func (cli *ETHCli) ORC721TokenURI(token string, tokenId *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721URIStorageAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("tokenURI", tokenId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return "", err
	}

	results, err := ins.Unpack("tokenURI", bz)
	if err != nil {
		return "", err
	}

	return results[0].(string), nil
}
