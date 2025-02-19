package ethcli

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// GenKey
// return pk,address,sk
func GenKey() (string, string, string, error) {
	privkey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", "", err
	}

	buff := make([]byte, 32)
	copy(buff[32-len(privkey.D.Bytes()):], privkey.D.Bytes())
	return common.Bytes2Hex(crypto.CompressPubkey(&privkey.PublicKey)),
		crypto.PubkeyToAddress(privkey.PublicKey).String(),
		common.Bytes2Hex(buff),
		nil
}

func HexToBytes(str string) []byte {
	str = strings.TrimPrefix(str, "0x")
	b, _ := hex.DecodeString(str)
	return b
}

func BytesToHex(bz []byte) string {
	return hex.EncodeToString(bz)
}

func BytesToHexWith0x(bz []byte) string {
	return "0x" + hex.EncodeToString(bz)
}

func ValidAddress(address string) bool {
	return common.IsHexAddress(address)
}

func ToWei(v *big.Int) *big.Int {
	return new(big.Int).Mul(v, big.NewInt(1e10))
}

func ToEther(v *big.Int) *big.Int {
	return new(big.Int).Div(v, big.NewInt(1e10))
}

func HashToAddress(hx common.Hash) common.Address {
	a := common.Address{}
	a.SetBytes(hx.Bytes())
	return a
}

func HashToBigInt(hx common.Hash) *big.Int {
	v := new(big.Int)
	v.SetBytes(hx.Bytes())
	return v
}
