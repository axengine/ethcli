package ethcli

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func TestETHCli_ERC721MintWithTokenURI(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721MintWithTokenURI(exampleERC721Token, exampleERC721IssuerKey, exampleFromAddress, "1001,0,1001")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721DecodeTokenId(t *testing.T) {
	decodeTokenId("0x7784f59caa8165c004cc6d4092c12ac1769568096c9b95909a9aa418db91a47b")
}

func decodeTokenId(hash string) (*big.Int, error) {
	cli, _ := New(exampleRawHTTPUrl)
	receipt, err := cli.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return nil, err
	}

	for _, v := range receipt.Logs {
		if len(v.Topics) == 0 {
			return nil, errors.New("No Topic")
		}
		ins, _ := abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"}]`))
		ev, err := ins.EventByID(v.Topics[0])
		if err != nil {
			continue
		}
		if ev.Name != "Transfer" {
			continue
		}

		from := hashToAddress(v.Topics[1])
		to := hashToAddress(v.Topics[2])
		tokenId := hashToBigInt(v.Topics[3])
		fmt.Printf("from:%s to:%s tokenId:%v\n", from.Hex(), to.Hex(), tokenId.Int64())
		return tokenId, nil
	}
	return nil, errors.New("no ERC721 Transfer")
}

func hashToAddress(hx common.Hash) common.Address {
	a := common.Address{}
	a.SetBytes(hx.Bytes())
	return a
}

func hashToBigInt(hx common.Hash) *big.Int {
	v := new(big.Int)
	v.SetBytes(hx.Bytes())
	return v
}

func TestETHCli_ERC721Burn(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721Burn(exampleERC721Token, exampleFromKey, big.NewInt(67))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721BalanceOf(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721BalanceOf("0x0c6F5145390028B7d09b513f47893b10C4c3a457", "0x69cB74cB86A05bd36d53Ccef1af3e90a6C66dD30", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721OwnerOf(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721OwnerOf(exampleERC721Token, big.NewInt(71), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721TokenOfOwnerByIndex(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721TokenOfOwnerByIndex(exampleERC721Token, exampleFromAddress, big.NewInt(0), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721TotalSupply(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721TotalSupply(exampleERC721Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721TokenByIndex(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721TokenByIndex(exampleERC721Token, big.NewInt(0), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC20Name(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC20Name(exampleERC721Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721Symbol(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721Symbol(exampleERC721Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721TokenURI(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721TokenURI(exampleERC721Token, big.NewInt(66), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721Paused(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721Paused(exampleERC721Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721Pause(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721Pause(exampleERC721Token, exampleERC721IssuerKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721Unpause(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721Unpause(exampleERC721Token, exampleERC721IssuerKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721TransferFrom(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721TransferFrom(exampleERC721Token, exampleFromKey, exampleFromAddress, exampleToAddress, big.NewInt(69))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721SafeTransferFrom(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721SafeTransferFrom(exampleERC721Token, exampleFromKey, exampleFromAddress, exampleERC721Token, big.NewInt(71))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721Approve(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721Approve(exampleERC721Token, exampleFromKey, exampleToAddress, big.NewInt(70))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721GetApproved(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721GetApproved(exampleERC721Token, big.NewInt(70), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721SetApprovalForAll(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721SetApprovalForAll(exampleERC721Token, exampleFromKey, exampleToAddress, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721IsApprovedForAll(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC721IsApprovedForAll(exampleERC721Token, exampleFromAddress, exampleToAddress, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC721Supports(t *testing.T) {
	cli, _ := New("https://bsc-dataseed1.binance.org:443")
	result, err := cli.ERC721SupportsInterface("0x6D07C33ad397d73Cbd3fE5349eF223ed36FE5b28", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
