package ethcli

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

var openzeppelinERC20MintBurnAbleAbi = `[{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"burn","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"burnFrom","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

func (cli *ETHCli) ORC20Mint(token, key, to, value string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20MintBurnAbleAbi))
	if err != nil {
		return "", err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", errors.New("invalid value:" + value)
	}
	data, _ := ins.Pack("mint", common.HexToAddress(to), amount)

	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC20Burn(token, key, value string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20MintBurnAbleAbi))
	if err != nil {
		return "", err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", errors.New("invalid value:" + value)
	}
	data, _ := ins.Pack("burn", amount)

	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC20BurnFrom(token, key, owner, value string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20MintBurnAbleAbi))
	if err != nil {
		return "", err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", errors.New("invalid value:" + value)
	}
	data, _ := ins.Pack("burnFrom", common.HexToAddress(owner), amount)

	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}
