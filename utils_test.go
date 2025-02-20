package ethcli

import (
	"fmt"
	"testing"
)

func TestGenKey(t *testing.T) {
	pk, address, sk, _ := GenKey()
	fmt.Printf("SK:%s PK:%s Address:%s\n", sk, pk, address)
}
