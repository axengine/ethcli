package ethcli

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

var (
	openzeppelinIERC721Abi            = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"approved","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"approve","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"balance","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"internalType":"address","name":"operator","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"owner","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"_approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"transferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	customERC721Exists                = `[{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"exists","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
	customERC721Mint                  = `[{"inputs":[{"internalType":"address","name":"to","type":"address"}],"name":"mint","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"}]`
	customERC721MintWithURI           = `[{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"string","name":"_tokenURI","type":"string"}],"name":"mint","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"}]`
	customERC721MintWithTokenIdAndURI = `[{"inputs":[{"internalType":"address","name":"receiver","type":"address"},{"internalType":"uint256","name":"_tokenId","type":"uint256"},{"internalType":"string","name":"_tokenURI","type":"string"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	customERC721SupportsInterface     = `[{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
)

func (cli *ETHCli) ORC721BalanceOf(token string, owner string, blockNumber *big.Int) (*big.Int, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC721Abi))
	if err != nil {
		return nil, err
	}
	data, _ := ins.Pack("balanceOf", common.HexToAddress(owner))

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

func (cli *ETHCli) ORC721OwnerOf(token string, tokenId *big.Int, blockNumber *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC721Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("ownerOf", tokenId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return "", err
	}

	results, err := ins.Unpack("ownerOf", bz)
	if err != nil {
		return "", err
	}

	return results[0].(common.Address).Hex(), nil
}

func (cli *ETHCli) ORC721SafeTransferFrom(token string, key, from, to string, tokenId *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC721Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("safeTransferFrom", common.HexToAddress(from), common.HexToAddress(to), tokenId)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC721TransferFrom(token string, key, from, to string, tokenId *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC721Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("transferFrom", common.HexToAddress(from), common.HexToAddress(to), tokenId)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC721Approve(token string, key, to string, tokenId *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC721Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("approve", common.HexToAddress(to), tokenId)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC721GetApproved(token string, tokenId *big.Int, blockNumber *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC721Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("getApproved", tokenId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return "", err
	}

	results, err := ins.Unpack("getApproved", bz)
	if err != nil {
		return "", err
	}

	return results[0].(common.Address).Hex(), nil
}

func (cli *ETHCli) ORC721SetApprovalForAll(token string, key, operator string, approved bool) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC721Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("setApprovalForAll", common.HexToAddress(operator), approved)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC721IsApprovedForAll(token string, owner, operator string, blockNumber *big.Int) (bool, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC721Abi))
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

func (cli *ETHCli) ORC721SafeTransferFromWithData(token string, key, from, to string, tokenId *big.Int, calldata []byte) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinIERC721Abi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("safeTransferFrom", common.HexToAddress(from), common.HexToAddress(to), tokenId, calldata)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC721Mint(token string, key string, to string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(customERC721Mint))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("mint", common.HexToAddress(to))
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC721MintWithTokenURI(token string, key string, to string, uri string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(customERC721MintWithURI))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("mint", common.HexToAddress(to), uri)
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC721MintWithTokenIdAndURI(token string, key string, to string, tokenId *big.Int, uri string) (string, error) {
	ins, err := abi.JSON(strings.NewReader(customERC721MintWithTokenIdAndURI))
	if err != nil {
		return "", err
	}
	data, err := ins.Pack("mint", common.HexToAddress(to), tokenId, uri)
	if err != nil {
		return "", err
	}
	return cli.SendMondoTx(key, &token, "0", BytesToHex(data), "0", 0)
}

func (cli *ETHCli) ORC721Exists(token string, tokenId *big.Int, blockNumber *big.Int) (bool, error) {
	ins, err := abi.JSON(strings.NewReader(customERC721Exists))
	if err != nil {
		return false, err
	}
	data, _ := ins.Pack("exists", tokenId)

	contract := common.HexToAddress(token)
	bz, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, blockNumber)
	if err != nil {
		return false, err
	}

	results, err := ins.Unpack("exists", bz)
	if err != nil {
		return false, err
	}

	return results[0].(bool), nil
}

func (cli *ETHCli) ORC721SupportsInterface(token string, blockNumber *big.Int) (bool, error) {
	ins, err := abi.JSON(strings.NewReader(customERC721SupportsInterface))
	if err != nil {
		return false, err
	}
	const _InterfaceId_ERC721 = "0x80ac58cd"
	interfaceIdBz, _ := hexutil.Decode(_InterfaceId_ERC721)
	var bs [4]byte
	copy(bs[:], interfaceIdBz)
	data, err := ins.Pack("supportsInterface", bs)
	if err != nil {
		return false, err
	}
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
