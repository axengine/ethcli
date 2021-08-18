package ethcli

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// BuildTx 生成交易
// nonce:from 的nonce,通过PendingNonceAt查询
// gasPrice 小数位18位，最小为1*1e10，可通过SuggestGasPrice查询建议值
// gas OLO普通转账21000,合约执行可通过EstimateGas先估算
// to nil时部署合约
// value OLO金额，小数位18位，最小为1*1e10
// data 合约执行字节码
func (cli *ETHCli) BuildTx(nonce uint64, gasPrice *big.Int,
	gas uint64, to *common.Address, value *big.Int, data []byte) *types.Transaction {
	return types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gas,
		To:       to,
		Value:    value,
		Data:     data,
	})
}

// SignTx 交易签名
// tx 待签名交易
// chainID 链id，Mondo(devnet:386 testnet:8724 mainnet:8723)
// key 私钥，无0x开头的hex格式
// 返回已签名tx，tx.Hash返回交易hash
func (cli *ETHCli) SignTx(tx *types.Transaction, chainID *big.Int, key string) (*types.Transaction, error) {
	priKey, _ := crypto.HexToECDSA(key)
	return cli.signTx(tx, chainID, priKey)
}

func (cli *ETHCli) signTx(tx *types.Transaction, chainID *big.Int, key *ecdsa.PrivateKey) (*types.Transaction, error) {
	signer := types.NewEIP155Signer(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
	if err != nil {
		return tx, err
	}
	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		return tx, err
	}
	return signedTx, nil
}

// SendTx 发送已签名交易
func (cli *ETHCli) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return cli.SendTransaction(ctx, signedTx)
}

// SendMondoTx 高级别的发送Mondo交易
// key 私钥，无0x开头的hex格式
// to 地址，为空或全0时为部署合约
// amount 金额，十进制字符串，整数，小数位18，第9-18位小数将被忽略
// payload 合约负载，0x开头、非0x开头的hex格式
// gasPrice 可选，金额，十进制字符串，整数，小数位18，第9-18位小数将被忽略;为0或不填时从链上时建议price
// gasLimit 可选，为0时实时估算
func (cli *ETHCli) SendMondoTx(key string, to *string, amount string, payload string, gasPrice string, gasLimit uint64) (string, error) {
	priKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return "", err
	}
	from := crypto.PubkeyToAddress(priKey.PublicKey)
	nonce, err := cli.PendingNonceAt(context.Background(), from)
	if err != nil {
		return "", err
	}

	var toAddr *common.Address
	if to == nil {
		toAddr = nil
	} else {
		tmp := common.HexToAddress(*to)
		toAddr = &tmp
	}

	value, _ := new(big.Int).SetString(amount, 10)
	if value == nil {
		value = big.NewInt(0)
	}

	data := HexToBytes(payload)

	price, _ := new(big.Int).SetString(gasPrice, 10)
	if price == nil || price.Cmp(big.NewInt(0)) == 0 {
		price, err = cli.SuggestGasPrice(context.Background())
		if err != nil {
			return "", err
		}
	}

	if gasLimit == 0 {
		gasLimit, err = cli.EstimateGas(context.Background(), ethereum.CallMsg{
			From:     from,
			To:       toAddr,
			GasPrice: price,
			Value:    value,
			Data:     data,
		})
		if err != nil {
			return "", err
		}
	}

	tx := cli.BuildTx(nonce, price, gasLimit, toAddr, value, data)
	signedTx, err := cli.signTx(tx, cli.chainID(), priKey)
	if err != nil {
		return "", err
	}

	err = cli.SendTransaction(context.Background(), signedTx)

	return signedTx.Hash().Hex(), err
}
