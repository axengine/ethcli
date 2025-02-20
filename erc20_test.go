package ethcli

import (
	"context"
	"fmt"
	"testing"
)

func Test_ERC20Name(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	name, err := cli.ERC20Name(context.Background(), exampleERC20Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(name)
}

func Test_ERC20Symbol(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	symbol, err := cli.ERC20Symbol(context.Background(), exampleERC20Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(symbol)
}

func Test_ERC20Decimals(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	decimals, err := cli.ERC20Decimals(context.Background(), exampleERC20Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decimals)
}

func Test_ERC20TotalSupply(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	totalSupply, err := cli.ERC20TotalSupply(context.Background(), exampleERC20Token, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(totalSupply)
}

func Test_ERC20BalanceOf(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	bal, err := cli.ERC20BalanceOf(context.Background(), exampleERC20Token,
		exampleToAddress, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(bal)
}

func Test_ERC20Transfer(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC20Transfer(context.Background(), exampleERC20Token, exampleFromKey, exampleToAddress, "1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_ERC20Approve(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC20Approve(context.Background(), exampleERC20Token, exampleFromKey, exampleToAddress, "100")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_ERC20Allowance(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC20Allowance(context.Background(), exampleERC20Token, exampleFromAddress, exampleToAddress, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_ERC20TransferFrom(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ERC20TransferFrom(context.Background(), exampleERC20Token, exampleFromKey, exampleFromAddress, exampleToAddress, "100")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
