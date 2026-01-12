package ethcli

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BuildLegacyTx build Legacy transaction
func BuildLegacyTx(nonce uint64, gasPrice *big.Int,
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

// BuildDynamicFeeTx build DynamicFee transaction
func BuildDynamicFeeTx(chainId *big.Int, nonce uint64, baseFee, priorityFee *big.Int,
	gas uint64, to *common.Address, value *big.Int, data []byte) *types.Transaction {
	dynamicFeeTx := &types.DynamicFeeTx{
		ChainID:    chainId,
		Nonce:      nonce,
		GasTipCap:  priorityFee,
		GasFeeCap:  new(big.Int).Add(baseFee, priorityFee),
		Gas:        gas,
		To:         to,
		Value:      value,
		Data:       data,
		AccessList: nil,
	}
	return types.NewTx(dynamicFeeTx)
}

// SignTx signs a transaction with the given private key.
func SignTx(tx *types.Transaction, chainId *big.Int, key *ecdsa.PrivateKey) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(chainId)
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

// SendTx Send signed transaction
func SendTx(ctx context.Context, cli *ethclient.Client, signedTx *types.Transaction) error {
	return cli.SendTransaction(ctx, signedTx)
}

// SendLegacyTx High-level Send Legacy Transaction
func SendLegacyTx(ctx context.Context, cli *ethclient.Client, key string, to *string, amount string, payload string, gasPrice string, gasLimit uint64) (string, error) {
	priKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return "", err
	}
	from := crypto.PubkeyToAddress(priKey.PublicKey)
	nonce, err := cli.PendingNonceAt(ctx, from)
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
		gasLimit, err = cli.EstimateGas(ctx, ethereum.CallMsg{
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

	tx := BuildLegacyTx(nonce, price, gasLimit, toAddr, value, data)

	chainId, err := cli.ChainID(ctx)
	if err != nil {
		return "", err
	}
	signedTx, err := SignTx(tx, chainId, priKey)
	if err != nil {
		return "", err
	}

	err = cli.SendTransaction(ctx, signedTx)

	return signedTx.Hash().Hex(), err
}

// SendDynamicFeeTx High-level Send DynamicFee Transaction
func SendDynamicFeeTx(ctx context.Context, cli *ethclient.Client, key string, to *string, amount string, payload string) (string, error) {
	priKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return "", err
	}
	from := crypto.PubkeyToAddress(priKey.PublicKey)
	nonce, err := cli.PendingNonceAt(ctx, from)
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

	block, err := cli.BlockByNumber(ctx, nil)
	if err != nil {
		return "", err
	}
	baseFee := block.BaseFee()
	priorityFeePerGas, err := cli.SuggestGasTipCap(context.Background())
	if err != nil {
		return "", err
	}

	chainId, err := cli.ChainID(ctx)
	if err != nil {
		return "", err
	}

	gas, err := cli.EstimateGas(ctx, ethereum.CallMsg{
		From:  from,
		To:    toAddr,
		Value: value,
		Data:  data,
	})
	if err != nil {
		return "", err
	}
	tx := BuildDynamicFeeTx(chainId, nonce, baseFee, priorityFeePerGas, gas, toAddr, value, data)

	signedTx, err := SignTx(tx, chainId, priKey)
	if err != nil {
		return "", err
	}

	err = cli.SendTransaction(ctx, signedTx)

	return signedTx.Hash().Hex(), err
}
