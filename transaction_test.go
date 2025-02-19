package ethcli

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestETHCli_SendOfflineTransaction(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	to := common.HexToAddress(exampleToAddress)

	nonce, err := cli.PendingNonceAt(context.Background(), common.HexToAddress(exampleFromAddress))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce=", nonce)
	tx := cli.BuildLegacyTx(nonce,
		ToWei(big.NewInt(100)),
		21000,
		&to,
		ToWei(big.NewInt(11*1e7)), // 1.1 ether
		nil)
	signedtx, err := cli.SignLegacyTx(tx, exampleFromKey)
	if err != nil {
		t.Fatal(err)
	}
	hash := signedtx.Hash()
	fmt.Println("hash=", hash.Hex())
	if err := cli.SendTx(context.Background(), signedtx); err != nil {
		t.Fatal(err)
	}
}

func TestETHCli_SendLegacyTx(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	to := exampleToAddress
	hash, err := cli.SendLegacyTx(exampleFromKey,
		&to,
		ToWei(big.NewInt(1)).String(), //0.00000001 ether
		"",
		ToWei(big.NewInt(100)).String(), // 0.000001 ether
		21000)
	if err != nil {
		t.Fatal(err, hash)
	}
	fmt.Println("hash=", hash)
}

func TestETHCli_SendLegacyTxWithoutFee(t *testing.T) {
	cli, _ := New(exampleRawHTTPUrl)
	to := exampleToAddress
	hash, err := cli.SendLegacyTx(exampleFromKey,
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
