package ethcli

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

var (
	openzeppelinERC721MetadataAbi = `[{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
)

func (cli *ETHCli) ORC721Name(token string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721MetadataAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("name")

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return "", err
	}

	results, err := ins.Unpack("name", bz)
	if err != nil {
		return "", err
	}

	return results[0].(string), nil
}

func (cli *ETHCli) ORC721Symbol(token string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721MetadataAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("symbol")

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return "", err
	}

	results, err := ins.Unpack("symbol", bz)
	if err != nil {
		return "", err
	}

	return results[0].(string), nil
}

// ORC721TokenURI
// for ERC721Metadata && ERC721URIStorage
//func (cli *ETHCli) ORC721TokenURI(token string, tokenId *big.Int) (string, error)
