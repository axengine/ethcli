package ethcli

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
	"testing"
)

var (
	exampleRawurl      = "http://192.168.10.106:8545"
	exampleChainId     = int64(386)
	exampleERC20Token  = "0x67EbBA731DCd5b763F5699650920a637eDbBEb93"
	exampleFromAddress = "0xED62EcAa9f0A43a92eA8ad08F199DF88Fc582F44"
	exampleFromKey     = "1259fe6a12097aa514d595dbc4a659d97d3fce7ae27d355c59117aa91c699c15"
	exampleToAddress   = "0x0F508F143E77b39F8e20DD9d2C1e515f0f527D9F"

	exampleERC721Token = "0x2A8eb9802a7d692128E2EDbfBcda2E71DCd74c01"
)

func Test_ChainID(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.ChainID(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_BlockByHash(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.BlockByHash(context.Background(), common.HexToHash("5FA8D136A49551A82F7DE7BB41C9B3EA58B3B8B468CC7E12DE811A080D0B8400"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_BlockByNumber(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.BlockByNumber(context.Background(), big.NewInt(9867045))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_BlockNumber(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_HeaderByHash(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.HeaderByHash(context.Background(), common.HexToHash("5FA8D136A49551A82F7DE7BB41C9B3EA58B3B8B468CC7E12DE811A080D0B8400"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_HeaderByNumber(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.HeaderByNumber(context.Background(), big.NewInt(9867045))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_TransactionByHash(t *testing.T) {
	cli := New(exampleRawurl)
	result, pending, err := cli.TransactionByHash(context.Background(), common.HexToHash("0x0f195566d383b3ee9174f0ea529bc604cb3e2d10bce8326c0a7eb05d2d347708"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", pending)
}

func Test_TransactionCount(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.TransactionCount(context.Background(), common.HexToHash("B725733BEB341B05E6C6E277CD20A158F8678CBC04C0925F33CA9104979B8EF8"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_TransactionInBlock(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.TransactionInBlock(context.Background(), common.HexToHash("B725733BEB341B05E6C6E277CD20A158F8678CBC04C0925F33CA9104979B8EF8"), 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_TransactionReceipt(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.TransactionReceipt(context.Background(), common.HexToHash("0xb3cd375f353433709221b78f58dcd8d474bc1176adf40cb1785021857f176c12"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

// unsupported method
func Test_SyncProgress(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.SyncProgress(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

// notifications not supported
func Test_SubscribeNewHead(t *testing.T) {
	cli := New(exampleRawurl)

	ch := make(chan *types.Header, 1)

	sub, err := cli.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", sub)
	defer sub.Unsubscribe()

	select {
	case h := <-ch:
		fmt.Printf("Got new header:%+v\n", h)
	}

}

func Test_NetworkID(t *testing.T) {
	cli := New(exampleRawurl)
	networkId, err := cli.NetworkID(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(networkId)
}

func Test_BalanceAt(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.BalanceAt(context.Background(), common.HexToAddress(exampleFromAddress), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_StorageAt(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.StorageAt(context.Background(), common.HexToAddress(exampleFromAddress), common.HexToHash(""), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_CodeAt(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.CodeAt(context.Background(), common.HexToAddress(exampleERC20Token), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_NonceAt(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.NonceAt(context.Background(), common.HexToAddress(exampleFromAddress), big.NewInt(9866995))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_FilterLogs(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: nil,
		Topics:    nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_SubscribeFilterLogs(t *testing.T) {
	cli := New(exampleRawurl)
	ch := make(chan types.Log)
	sub, err := cli.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: nil,
		Topics:    nil,
	}, ch)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", sub)

	defer sub.Unsubscribe()

	select {
	case h := <-ch:
		fmt.Printf("Got new log:%+v\n", h)
	}
}

func Test_PendingBalanceAt(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.PendingBalanceAt(context.Background(), common.HexToAddress(exampleFromAddress))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_PendingStorageAt(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.PendingStorageAt(context.Background(), common.HexToAddress(exampleFromAddress), common.HexToHash(""))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_PendingNonceAt(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.PendingNonceAt(context.Background(), common.HexToAddress(exampleFromAddress))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_PendingTransactionCount(t *testing.T) {
	cli := New(exampleRawurl)
	result, err := cli.PendingTransactionCount(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_CallContract(t *testing.T) {
	cli := New(exampleRawurl)

	contract := common.HexToAddress(exampleERC20Token)

	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
	if err != nil {
		t.Fatal(err)
	}
	data, _ := ins.Pack("name")

	result, err := cli.CallContract(context.Background(),
		ethereum.CallMsg{
			To:   &contract,
			Data: data,
		}, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_PendingCallContract(t *testing.T) {
	cli := New(exampleRawurl)

	contract := common.HexToAddress(exampleERC20Token)

	ins, err := abi.JSON(strings.NewReader(openzeppelinERC20Abi))
	if err != nil {
		t.Fatal(err)
	}
	data, _ := ins.Pack("name")

	result, err := cli.PendingCallContract(context.Background(),
		ethereum.CallMsg{
			To:   &contract,
			Data: data,
		})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_SuggestGasPrice(t *testing.T) {
	cli := New(exampleRawurl)

	result, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func Test_EstimateGas(t *testing.T) {
	cli := New(exampleRawurl)
	contract := common.HexToAddress(exampleToAddress)
	result, err := cli.EstimateGas(context.Background(),
		ethereum.CallMsg{
			To:    &contract,
			Value: big.NewInt(0),
			Data:  nil,
		})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}
