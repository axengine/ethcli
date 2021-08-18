package ethcli

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
)

// GenKey 本地生成mondo链格式的账户
// 返回：公钥，地址，私钥，mondo采用压缩公钥
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
