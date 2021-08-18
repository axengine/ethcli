package ethcli

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

var openzeppelinERC20Abi = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_who","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_name","type":"string"},{"name":"_symbol","type":"string"},{"name":"_decimals","type":"uint8"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`

func (cli *ETHCli) ORC20Name(token string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
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

func (cli *ETHCli) ORC20Symbol(token string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
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

func (cli *ETHCli) ORC20Decimals(token string) (uint8, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
	if err != nil {
		return 0, err
	}
	data, _ := ins.Pack("decimals")

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return 0, err
	}

	results, err := ins.Unpack("decimals", bz)
	if err != nil {
		return 0, err
	}

	return results[0].(uint8), nil
}

func (cli *ETHCli) ORC20TotalSupply(token string) (*big.Int, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
	if err != nil {
		return nil, err
	}
	data, _ := ins.Pack("totalSupply")

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	results, err := ins.Unpack("totalSupply", bz)
	if err != nil {
		return nil, err
	}

	return results[0].(*big.Int), nil
}

func (cli *ETHCli) ORC20BalanceOf(token string, address string) (*big.Int, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
	if err != nil {
		return nil, err
	}
	data, _ := ins.Pack("balanceOf", common.HexToAddress(address))

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	results, err := ins.Unpack("balanceOf", bz)
	if err != nil {
		return nil, err
	}

	return results[0].(*big.Int), nil
}

func (cli *ETHCli) ORC20Transfer(key, token, to, value string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
	if err != nil {
		return "", err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", errors.New("invalid value:" + value)
	}
	data, _ := ins.Pack("transfer", common.HexToAddress(to), amount)

	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC20Allowance(token, owner, spender string) (*big.Int, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
	if err != nil {
		return nil, err
	}
	data, _ := ins.Pack("allowance", common.HexToAddress(owner), common.HexToAddress(spender))

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	results, err := ins.Unpack("allowance", bz)
	if err != nil {
		return nil, err
	}

	return results[0].(*big.Int), nil
}

func (cli *ETHCli) ORC20TransferFrom(key, token, from, to, value string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
	if err != nil {
		return "", err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", errors.New("invalid value:" + value)
	}
	data, _ := ins.Pack("transferFrom", common.HexToAddress(from), common.HexToAddress(to), amount)

	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC20Approve(key, token, spender, value string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
	if err != nil {
		return "", err
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", errors.New("invalid value:" + value)
	}
	data, _ := ins.Pack("approve", common.HexToAddress(spender), amount)

	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}
