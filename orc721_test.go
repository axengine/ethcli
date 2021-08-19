package ethcli

import (
	"fmt"
	"math/big"
	"testing"
)

func TestETHCli_ORC721MintWithTokenURI(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721MintWithTokenURI(exampleERC721Token, exampleERC721IssuerKey, exampleFromAddress, "1001,0,1001")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Burn(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Burn(exampleERC721Token, exampleFromKey, big.NewInt(67))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721BalanceOf(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721BalanceOf(exampleERC721Token, exampleFromAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721OwnerOf(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721OwnerOf(exampleERC721Token, big.NewInt(71))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721TokenOfOwnerByIndex(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721TokenOfOwnerByIndex(exampleERC721Token, exampleFromAddress, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721TotalSupply(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721TotalSupply(exampleERC721Token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721TokenByIndex(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721TokenByIndex(exampleERC721Token, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC20Name(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC20Name(exampleERC721Token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Symbol(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Symbol(exampleERC721Token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721TokenURI(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721TokenURI(exampleERC721Token, big.NewInt(66))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Paused(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Paused(exampleERC721Token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Pause(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Pause(exampleERC721Token, exampleERC721IssuerKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Unpause(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Unpause(exampleERC721Token, exampleERC721IssuerKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721TransferFrom(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721TransferFrom(exampleERC721Token, exampleFromKey, exampleFromAddress, exampleToAddress, big.NewInt(69))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721SafeTransferFrom(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721SafeTransferFrom(exampleERC721Token, exampleFromKey, exampleFromAddress, exampleERC721Token, big.NewInt(71))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721Approve(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721Approve(exampleERC721Token, exampleFromKey, exampleToAddress, big.NewInt(70))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721GetApproved(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721GetApproved(exampleERC721Token, big.NewInt(70))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721SetApprovalForAll(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721SetApprovalForAll(exampleERC721Token, exampleFromKey, exampleToAddress, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ORC721IsApprovedForAll(t *testing.T) {
	cli := New(exampleRawHTTPUrl)
	result, err := cli.ORC721IsApprovedForAll(exampleERC721Token, exampleFromAddress, exampleToAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
