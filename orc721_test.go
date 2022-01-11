package ethcli

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"testing"
)

func TestETHCli_ORC721MintWithTokenURI(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721MintWithTokenURI(exampleERC721Token, exampleERC721IssuerKey, exampleFromAddress, "1001,0,1001")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721DecodeTokenId(t *testing.T) {
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

func TestETHCli_ORC721Burn(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Burn(exampleERC721Token, exampleFromKey, big.NewInt(67))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721BalanceOf(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721BalanceOf("0x0c6F5145390028B7d09b513f47893b10C4c3a457", "0x69cB74cB86A05bd36d53Ccef1af3e90a6C66dD30", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721OwnerOf(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721OwnerOf(exampleERC721Token, big.NewInt(71), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721TokenOfOwnerByIndex(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721TokenOfOwnerByIndex(exampleERC721Token, exampleFromAddress, big.NewInt(0), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721TotalSupply(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721TotalSupply(exampleERC721Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721TokenByIndex(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721TokenByIndex(exampleERC721Token, big.NewInt(0), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC20Name(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC20Name(exampleERC721Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Symbol(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Symbol(exampleERC721Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721TokenURI(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721TokenURI(exampleERC721Token, big.NewInt(66), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Paused(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Paused(exampleERC721Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Pause(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Pause(exampleERC721Token, exampleERC721IssuerKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Unpause(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Unpause(exampleERC721Token, exampleERC721IssuerKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721TransferFrom(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721TransferFrom(exampleERC721Token, exampleFromKey, exampleFromAddress, exampleToAddress, big.NewInt(69))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721SafeTransferFrom(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721SafeTransferFrom(exampleERC721Token, exampleFromKey, exampleFromAddress, exampleERC721Token, big.NewInt(71))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Approve(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Approve(exampleERC721Token, exampleFromKey, exampleToAddress, big.NewInt(70))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721GetApproved(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721GetApproved(exampleERC721Token, big.NewInt(70), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721SetApprovalForAll(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721SetApprovalForAll(exampleERC721Token, exampleFromKey, exampleToAddress, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721IsApprovedForAll(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC721IsApprovedForAll(exampleERC721Token, exampleFromAddress, exampleToAddress, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Supports(t *testing.T) {
	cli, _ := New("https://bsc-dataseed1.binance.org:443")
	result, err := cli.ORC721SupportsInterface("0x6D07C33ad397d73Cbd3fE5349eF223ed36FE5b28", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
