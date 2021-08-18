package ethcli

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestETHCli_SendOfflineTransaction(t *testing.T) {
	cli := New(exampleRawurl)
	to := common.HexToAddress(exampleToAddress)

	nonce, err := cli.PendingNonceAt(context.Background(), common.HexToAddress(exampleFromAddress))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce=", nonce)
	tx := cli.BuildTx(nonce,
		ToWei(big.NewInt(100)),
		21000,
		&to,
		ToWei(big.NewInt(11*1e7)), // 1.1 OLO
		nil)
	signedtx, err := cli.SignTx(tx, big.NewInt(exampleChainId), exampleFromKey)
	if err != nil {
		t.Fatal(err)
	}
	hash := signedtx.Hash()
	fmt.Println("hash=", hash.Hex())
	if err := cli.SendTx(context.Background(), signedtx); err != nil {
		t.Fatal(err)
	}
}

func TestETHCli_SendMondoTx(t *testing.T) {
	cli := New(exampleRawurl)
	to := exampleToAddress
	hash, err := cli.SendMondoTx(exampleFromKey,
		&to,
		ToWei(big.NewInt(1)).String(), //0.00000001 OLO
		"",
		ToWei(big.NewInt(100)).String(), // 0.000001 OLO
		21000)
	if err != nil {
		t.Fatal(err, hash)
	}
	fmt.Println("hash=", hash)
}

func TestETHCli_SendMondoTxWithoutFee(t *testing.T) {
	cli := New(exampleRawurl)
	to := exampleToAddress
	hash, err := cli.SendMondoTx(exampleFromKey,
		&to,
		ToWei(big.NewInt(1)).String(), //0.00000001 OLO
		"",
		"",
		0)
	if err != nil {
		t.Fatal(err, hash)
	}
	fmt.Println("hash=", hash)
}
