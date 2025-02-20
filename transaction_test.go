package ethcli

import (
	"context"
	"fmt"
	"math/big"
	"testing"
)

func TestETHCli_SendLegacyTx(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	to := exampleToAddress
	hash, err := cli.SendLegacyTx(context.Background(), exampleFromKey,
		&to,
		ToWei(big.NewInt(1)).String(), //0.00000001 ether
		"",
		ToWei(big.NewInt(10)).String(), // 0.000001 ether
		21000)
	if err != nil {
		t.Fatal(err, hash)
	}
	fmt.Println("hash=", hash)
}

func TestETHCli_SendLegacyTxWithoutFee(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	to := exampleToAddress
	hash, err := cli.SendLegacyTx(context.Background(), exampleFromKey,
		&to,
		ToWei(big.NewInt(1)).String(), //0.00000001 ether
		"",
		"",
		0)
	if err != nil {
		t.Fatal(err, hash)
	}
	fmt.Println("hash=", hash)
}

func TestETHCli_SendDynamicFee(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	to := exampleToAddress
	hash, err := cli.SendDynamicFeeTx(context.Background(), exampleFromKey,
		&to,
		ToWei(big.NewInt(1)).String(), //0.00000001 ether
		"")
	if err != nil {
		t.Fatal(err, hash)
	}
	fmt.Println("hash=", hash)
}
