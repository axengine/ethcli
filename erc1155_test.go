package ethcli

import (
	"context"
	"fmt"
	"math/big"
	"testing"
)

var (
	bscTestnetRpc       = "https://data-seed-prebsc-1-s1.binance.org:8545"
	exampleERC1155Token = "0x2255444362b73087b82bb2c95239be4e708c130d"
)

func TestETHCli_ERC1155BalanceOf(t *testing.T) {
	cli, _ := New(bscTestnetRpc)
	result, err := cli.ERC1155BalanceOf(context.Background(), exampleERC1155Token, "0xf11804c522753e2afd2a4a8d9c1bf7ab0abaf60f", big.NewInt(128), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestETHCli_ERC1155Uri(t *testing.T) {
	cli, _ := New(bscTestnetRpc)
	result, err := cli.ERC1155Uri(context.Background(), exampleERC1155Token, big.NewInt(128), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
