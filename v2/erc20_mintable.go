package ethcli

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var openzeppelinERC20MintBurnAbleAbi = `[{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"burn","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"burnFrom","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

func ERC20Mint(ctx context.Context, cli *ethclient.Client, token, key, to, value string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20MintBurnAbleAbi))
	if err != nil {
		return "", err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", errors.New("invalid value:" + value)
	}
	data, _ := ins.Pack("mint", common.HexToAddress(to), amount)

	return SendLegacyTx(ctx, cli, key, &token, "0", BytesToHex(data), "0", 0)
}

func ERC20Burn(ctx context.Context, cli *ethclient.Client, token, key, value string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20MintBurnAbleAbi))
	if err != nil {
		return "", err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", errors.New("invalid value:" + value)
	}
	data, _ := ins.Pack("burn", amount)

	return SendLegacyTx(ctx, cli, key, &token, "0", BytesToHex(data), "0", 0)
}

func ERC20BurnFrom(ctx context.Context, cli *ethclient.Client, token, key, owner, value string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20MintBurnAbleAbi))
	if err != nil {
		return "", err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", errors.New("invalid value:" + value)
	}
	data, _ := ins.Pack("burnFrom", common.HexToAddress(owner), amount)

	return SendLegacyTx(ctx, cli, key, &token, "0", BytesToHex(data), "0", 0)
}

func ERC20MintData(to, value string) ([]byte, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20MintBurnAbleAbi))
	if err != nil {
		return nil, err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return nil, errors.New("invalid value:" + value)
	}
	return ins.Pack("mint", common.HexToAddress(to), amount)
}

func ERC20BurnData(value string) ([]byte, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20MintBurnAbleAbi))
	if err != nil {
		return nil, err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return nil, errors.New("invalid value:" + value)
	}
	return ins.Pack("burn", amount)
}

func ERC20BurnFromData(owner, value string) ([]byte, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20MintBurnAbleAbi))
	if err != nil {
		return nil, err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return nil, errors.New("invalid value:" + value)
	}
	return ins.Pack("burnFrom", common.HexToAddress(owner), amount)
}
