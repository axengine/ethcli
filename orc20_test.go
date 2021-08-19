package ethcli

import (
	"fmt"
	"testing"
)

func Test_ORC20Name(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	name, err := cli.ORC20Name(exampleERC20Token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(name)
}

func Test_ORC20Symbol(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	symbol, err := cli.ORC20Symbol(exampleERC20Token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(symbol)
}

func Test_ORC20Decimals(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	decimals, err := cli.ORC20Decimals(exampleERC20Token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decimals)
}

func Test_ORC20TotalSupply(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	totalSupply, err := cli.ORC20TotalSupply(exampleERC20Token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(totalSupply)
}

func Test_ORC20BalanceOf(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	bal, err := cli.ORC20BalanceOf(exampleERC20Token,
		exampleToAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(bal)
}

func Test_ORC20Transfer(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC20Transfer(exampleERC20Token, exampleFromKey, exampleToAddress, "1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_ORC20Approve(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC20Approve(exampleERC20Token, exampleFromKey, exampleToAddress, "100")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_ORC20Allowance(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC20Allowance(exampleERC20Token, exampleFromAddress, exampleToAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_ORC20TransferFrom(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	result, err := cli.ORC20TransferFrom(exampleERC20Token, exampleFromKey, exampleFromAddress, exampleToAddress, "100")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
