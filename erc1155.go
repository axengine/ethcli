package ethcli

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	openzeppelinIERC1155Abi = `[{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"burn","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"internalType":"uint256[]","name":"values","type":"uint256[]"}],"name":"burnBatch","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"internalType":"uint256[]","name":"amounts","type":"uint256[]"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"mintBatch","outputs":[],"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"indexed":false,"internalType":"uint256[]","name":"values","type":"uint256[]"}],"name":"TransferBatch","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"id","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"TransferSingle","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"value","type":"string"},{"indexed":true,"internalType":"uint256","name":"id","type":"uint256"}],"name":"URI","type":"event"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address[]","name":"accounts","type":"address[]"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"}],"name":"balanceOfBatch","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"internalType":"uint256[]","name":"amounts","type":"uint256[]"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeBatchTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"uri","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
)

func (cli *EvmClient) ERC1155BalanceOf(ctx context.Context, token string, owner string, tokenId *big.Int, blockNumber *big.Int) (*big.Int, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return nil, err
	}
	data, _ := ins.Pack("balanceOf", common.HexToAddress(owner), tokenId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
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

func (cli *EvmClient) ERC1155BalanceOfBatch(ctx context.Context, token string, owners []string, tokenIds []*big.Int, blockNumber *big.Int) ([]*big.Int, error) {
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
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
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

func (cli *EvmClient) ERC1155IsApprovedForAll(ctx context.Context, token string, owner string, operator string, blockNumber *big.Int) (bool, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return false, err
	}
	data, _ := ins.Pack("isApprovedForAll", common.HexToAddress(owner), common.HexToAddress(operator))

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
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

func (cli *EvmClient) ERC1155SafeBatchTransferFrom(ctx context.Context, key string, token string, owner string, to string, ids []*big.Int, amounts []*big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("safeBatchTransferFrom", common.HexToAddress(owner), common.HexToAddress(to), ids, amounts)
	return cli.SendLegacyTx(ctx, key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *EvmClient) ERC1155SafeTransferFrom(ctx context.Context, key string, token string, owner string, to string, id *big.Int, amount *big.Int, data []byte) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("safeTransferFrom", common.HexToAddress(owner), common.HexToAddress(to), id, amount, data)
	return cli.SendLegacyTx(ctx, key, &token, "0", BytesToHex(bz), "0", 0)
}

func (cli *EvmClient) ERC1155SetApprovalForAll(ctx context.Context, key string, token string, owner string, operator string, id *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("setApprovalForAll", common.HexToAddress(owner), common.HexToAddress(operator), id)
	return cli.SendLegacyTx(ctx, key, &token, "0", BytesToHex(bz), "0", 0)
}

func (cli *EvmClient) ERC1155SupportsInterface(ctx context.Context, token string, interfaceId [4]byte, blockNumber *big.Int) (bool, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return false, err
	}
	data, _ := ins.Pack("supportsInterface", interfaceId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
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

func (cli *EvmClient) ERC1155Uri(ctx context.Context, token string, tokenId *big.Int, blockNumber *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("uri", tokenId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(ctx, ethereum.CallMsg{
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

func (cli *EvmClient) ERC1155Mint(ctx context.Context, key string, token string, to string, id *big.Int, amount *big.Int, data []byte) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("mint", common.HexToAddress(to), id, amount, data)
	return cli.SendLegacyTx(ctx, key, &token, "0", BytesToHex(bz), "0", 0)
}

func (cli *EvmClient) ERC1155MintBatch(ctx context.Context, key string, token string, to string, ids []*big.Int, amounts []*big.Int, data []byte) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("mintBatch", common.HexToAddress(to), ids, amounts, data)
	return cli.SendLegacyTx(ctx, key, &token, "0", BytesToHex(bz), "0", 0)
}

func (cli *EvmClient) ERC1155Burn(ctx context.Context, key string, token string, to string, id *big.Int, amount *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("burn", common.HexToAddress(to), id, amount)
	return cli.SendLegacyTx(ctx, key, &token, "0", BytesToHex(bz), "0", 0)
}

func (cli *EvmClient) ERC1155BurnBatch(ctx context.Context, key string, token string, to string, ids []*big.Int, amounts []*big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC1155Abi))
	if err != nil {
		return "", err
	}
	bz, _ := ins.Pack("burnBatch", common.HexToAddress(to), ids, amounts)
	return cli.SendLegacyTx(ctx, key, &token, "0", BytesToHex(bz), "0", 0)
}
