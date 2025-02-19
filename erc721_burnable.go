package ethcli

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	openzeppelinERC721BurnableAbi = `[{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"burn","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
)

func (cli *EvmClient) ERC721Burn(token string, key string, tokenId *big.Int) (string, error) {
	ins, err := abi.JSON(strings.NewReader(openzeppelinERC721BurnableAbi))
	if err != nil {
		return "", err
	}
	data, _ := ins.Pack("burn", tokenId)
	return cli.SendLegacyTx(key, &token, "0", BytesToHex(data), "0", 0)
}
