package ethcli

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

func Test_sendTx(t *testing.T) {
	//client, err := ethclient.Dial("https://testnet-web3.wolot.io")
	client, err := ethclient.Dial("http://192.168.10.106:8545")
	if err != nil {
		panic(err)
	}

	key := "e8ca4b92b646487bf6be852e35dbe96496386a5541f16a7b33b84c96b5c2d0b0"
	//from := "0x35bECc25356cE434034ccCcfE986d4cD5109a75B"
	chainID := big.NewInt(386)

	to := common.HexToAddress("0x0F508F143E77b39F8e20DD9d2C1e515f0f527D9F")
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    1,
		GasPrice: big.NewInt(100 * 1e10),
		Gas:      21000,
		To:       &to,
		Value:    big.NewInt(1 * 1e10),
		Data:     []byte{},
	})

	signer := types.NewEIP155Signer(chainID)

	testKey, _ := crypto.HexToECDSA(key)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), testKey)
	if err != nil {
		t.Fatal(err)
	}
	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("hash:=", signedTx.Hash())

	//ctx,cancel:=context.WithTimeout(context.Background(),time.Second*5)

	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		t.Fatal(err)
	}
}

func Test_getNonce(t *testing.T) {
	//client, err := ethclient.Dial("https://testnet-web3.wolot.io")
	client, err := ethclient.Dial("http://192.168.10.106:8545")
	if err != nil {
		panic(err)
	}
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress("0x0F508F143E77b39F8e20DD9d2C1e515f0f527D9F"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(nonce)
}
