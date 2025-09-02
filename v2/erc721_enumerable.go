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
	openzeppelinERC721EnumerableAbi = `[{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"index","type":"uint256"}],"name":"tokenOfOwnerByIndex","outputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"index","type":"uint256"}],"name":"tokenByIndex","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
)

func ERC721TokenOfOwnerByIndex(ctx context.Context, cli *ethclient.Client, token string, owner string, index *big.Int, blockNumber *big.Int) (*big.Int, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721EnumerableAbi))
	if err != nil {
		return nil, err
	}
	data, _ := ins.Pack("tokenOfOwnerByIndex", common.HexToAddress(owner), index)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return nil, err
	}

	results, err := ins.Unpack("tokenOfOwnerByIndex", bz)
	if err != nil {
		return nil, err
	}

	return results[0].(*big.Int), nil
}

func ERC721TotalSupply(ctx context.Context, cli *ethclient.Client, token string, blockNumber *big.Int) (*big.Int, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721EnumerableAbi))
	if err != nil {
		return nil, err
	}
	data, _ := ins.Pack("totalSupply")

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return nil, err
	}

	results, err := ins.Unpack("totalSupply", bz)
	if err != nil {
		return nil, err
	}

	return results[0].(*big.Int), nil
}

func ERC721TokenByIndex(ctx context.Context, cli *ethclient.Client, token string, index *big.Int, blockNumber *big.Int) (*big.Int, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721EnumerableAbi))
	if err != nil {
		return nil, err
	}
	data, _ := ins.Pack("tokenByIndex", index)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return nil, err
	}

	results, err := ins.Unpack("tokenByIndex", bz)
	if err != nil {
		return nil, err
	}

	return results[0].(*big.Int), nil
}
