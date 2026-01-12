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
	openzeppelinERC721PauseableAbi = `[{"inputs":[],"name":"pause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"unpause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"paused","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
)

func ERC721Pause(ctx context.Context, cli *ethclient.Client, token string, key string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721PauseableAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("pause")
	return SendLegacyTx(ctx, cli, key, &token, "0", BytesToHex(data), "0", 0)
}

func ERC721Unpause(ctx context.Context, cli *ethclient.Client, token string, key string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721PauseableAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("unpause")
	return SendLegacyTx(ctx, cli, key, &token, "0", BytesToHex(data), "0", 0)
}

func ERC721Paused(ctx context.Context, cli *ethclient.Client, token string, blockNumber *big.Int) (bool, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721PauseableAbi))
	if err != nil {
		return false, err
	}
	data, _ := ins.Pack("paused")

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return false, err
	}

	results, err := ins.Unpack("paused", bz)
	if err != nil {
		return false, err
	}

	return results[0].(bool), nil
}

func ERC721PauseData() ([]byte, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721PauseableAbi))
	if err != nil {
		return nil, err
	}
	return ins.Pack("pause")
}

func ERC721UnpauseData() ([]byte, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721PauseableAbi))
	if err != nil {
		return nil, err
	}
	return ins.Pack("unpause")
}
