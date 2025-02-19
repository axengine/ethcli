package ethcli

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestGenKey(t *testing.T) {
	pk, address, sk, _ := GenKey()
	fmt.Printf("SK:%s PK:%s Address:%s\n", sk, pk, address)

	t.Log(common.HexToAddress("0x141238092b9e162f9798d25e2274ab4ccb28ca41"))
}
