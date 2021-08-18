package ethcli

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

var (
	openzeppelinERC721PauseableAbi = `[{"inputs":[],"name":"pause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"unpause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"paused","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
)

func (cli *ETHCli) ORC721Pause(token string, key string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721PauseableAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("pause")
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC721Unpause(token string, key string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721PauseableAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("unpause")
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC721Paused(token string) (bool, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721PauseableAbi))
	if err != nil {
		return false, err
	}
	data, _ := ins.Pack("paused")

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return false, err
	}

	results, err := ins.Unpack("paused", bz)
	if err != nil {
		return false, err
	}

	return results[0].(bool), nil
}
