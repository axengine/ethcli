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
	openzeppelinIERC1155Abi = `[{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"burn","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"internalType":"uint256[]","name":"values","type":"uint256[]"}],"name":"burnBatch","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"internalType":"uint256[]","name":"amounts","type":"uint256[]"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"mintBatch","outputs":[],"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"indexed":false,"internalType":"uint256[]","name":"values","type":"uint256[]"}],"name":"TransferBatch","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"id","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"TransferSingle","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"value","type":"string"},{"indexed":true,"internalType":"uint256","name":"id","type":"uint256"}],"name":"URI","type":"event"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address[]","name":"accounts","type":"address[]"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"}],"name":"balanceOfBatch","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"internalType":"uint256[]","name":"amounts","type":"uint256[]"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeBatchTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"uri","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
)

func (cli *ETHCli) ORC1155BalanceOf(token string, owner string, tokenId *big.Int, blockNumber *big.Int) (*big.Int, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return nil, err
	}
	data, _ := ins.Pack("balanceOf", common.HexToAddress(owner), tokenId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return nil, err
	}

	results, err := ins.Unpack("balanceOf", bz)
	if err != nil {
		return nil, err
	}

	return results[0].(*big.Int), nil
}

func (cli *ETHCli) ORC1155BalanceOfBatch(token string, owners []string, tokenIds []*big.Int, blockNumber *big.Int) ([]*big.Int, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return nil, err
	}
	var accounts []common.Address
	for _, v := range owners {
		accounts = append(accounts, common.HexToAddress(v))
	}
	data, _ := ins.Pack("balanceOfBatch", accounts, tokenIds)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return nil, err
	}

	results, err := ins.Unpack("balanceOfBatch", bz)
	if err != nil {
		return nil, err
	}

	return results[0].([]*big.Int), nil
}

func (cli *ETHCli) ORC1155IsApprovedForAll(token string, owner string, operator string, blockNumber *big.Int) (bool, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return false, err
	}
	data, _ := ins.Pack("isApprovedForAll", common.HexToAddress(owner), common.HexToAddress(operator))

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return false, err
	}

	results, err := ins.Unpack("isApprovedForAll", bz)
	if err != nil {
		return false, err
	}

	return results[0].(bool), nil
}

func (cli *ETHCli) ORC1155SafeBatchTransferFrom(key string, token string, owner string, to string, ids []*big.Int, amounts []*big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("safeBatchTransferFrom", common.HexToAddress(owner), common.HexToAddress(to), ids, amounts)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC1155SafeTransferFrom(key string, token string, owner string, to string, id *big.Int, amount *big.Int, data []byte) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("safeTransferFrom", common.HexToAddress(owner), common.HexToAddress(to), id, amount, data)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(bz), "0", 0)
}

func (cli *ETHCli) ORC1155SetApprovalForAll(key string, token string, owner string, operator string, id *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("setApprovalForAll", common.HexToAddress(owner), common.HexToAddress(operator), id)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(bz), "0", 0)
}

func (cli *ETHCli) ORC1155SupportsInterface(token string, interfaceId [4]byte, blockNumber *big.Int) (bool, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return false, err
	}
	data, _ := ins.Pack("supportsInterface", interfaceId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return false, err
	}

	results, err := ins.Unpack("supportsInterface", bz)
	if err != nil {
		return false, err
	}

	return results[0].(bool), nil
}

func (cli *ETHCli) ORC1155Uri(token string, tokenId *big.Int, blockNumber *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("uri", tokenId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return "", err
	}

	results, err := ins.Unpack("uri", bz)
	if err != nil {
		return "", err
	}

	return results[0].(string), nil
}

func (cli *ETHCli) ORC1155Mint(key string, token string, to string, id *big.Int, amount *big.Int, data []byte) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("mint", common.HexToAddress(to), id, amount, data)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(bz), "0", 0)
}

func (cli *ETHCli) ORC1155MintBatch(key string, token string, to string, ids []*big.Int, amounts []*big.Int, data []byte) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("mintBatch", common.HexToAddress(to), ids, amounts, data)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(bz), "0", 0)
}

func (cli *ETHCli) ORC1155Burn(key string, token string, to string, id *big.Int, amount *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("burn", common.HexToAddress(to), id, amount)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(bz), "0", 0)
}

func (cli *ETHCli) ORC1155BurnBatch(key string, token string, to string, ids []*big.Int, amounts []*big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("burnBatch", common.HexToAddress(to), ids, amounts)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(bz), "0", 0)
}
