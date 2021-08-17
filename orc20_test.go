package ethcli

import (
	"fmt"
	"testing"
)

func Test_ORC20Name(t *testing.T) {
	cli := New("http://192.168.10.106:8545")
	name, err := cli.ORC20Name("0x67EbBA731DCd5b763F5699650920a637eDbBEb93")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(name)
}

func Test_ORC20Symbol(t *testing.T) {
	cli := New("http://192.168.10.106:8545")
	symbol, err := cli.ORC20Symbol("0x67EbBA731DCd5b763F5699650920a637eDbBEb93")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(symbol)
}

func Test_ORC20Decimals(t *testing.T) {
	cli := New("http://192.168.10.106:8545")
	decimals, err := cli.ORC20Decimals("0x67EbBA731DCd5b763F5699650920a637eDbBEb93")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decimals)
}

func Test_ORC20TotalSupply(t *testing.T) {
	cli := New("http://192.168.10.106:8545")
	totalSupply, err := cli.ORC20TotalSupply("0x67EbBA731DCd5b763F5699650920a637eDbBEb93")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(totalSupply)
}

func Test_ORC20BalanceOf(t *testing.T) {
	cli := New("http://192.168.10.106:8545")
	bal, err := cli.ORC20BalanceOf("0x497cA24e4CC5898c842f9548fa04c3E4ED02d84b",
		"0xda1418b6C308ea5BB3987e3D5f3b47142bA048F8")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(bal)
}
