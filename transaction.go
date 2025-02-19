package ethcli

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// BuildLegacyTx generate transaction
// nonce: nonce from, query through PendingNonceAt
// gasPrice has 18 decimal places, the minimum is 1*1e10, the suggested value can be queried through SuggestGasPrice
// Gas ether ordinary transfer is 21,000, contract execution can be estimated first through EstimateGas
// Deploy the contract when to nil
// value ether amount, 18 decimal places, minimum 1*1e10
// data contract execution bytecode
func (cli *EvmClient) BuildLegacyTx(nonce uint64, gasPrice *big.Int,
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

// SignLegacyTx sign transaction
// tx transaction to be signed
// key private key, without 0x prefix
// Return signed tx, tx.Hash returns transaction hash
func (cli *EvmClient) SignLegacyTx(tx *types.Transaction, key string) (*types.Transaction, error) {
	priKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	chainId, err := cli.chainID()
	if err != nil {
		return nil, err
	}
	return cli.signTx(tx, chainId, priKey)
}

func (cli *EvmClient) signTx(tx *types.Transaction, chainId *big.Int, key *ecdsa.PrivateKey) (*types.Transaction, error) {
	signer := types.NewEIP155Signer(chainId)
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

// SendTx Send signed transactions
func (cli *EvmClient) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return cli.SendTransaction(ctx, signedTx)
}

// SendLegacyTx High-level Send Legacy Transactions
// key private key, no hex format starting from 0x
// to address, empty or all 0 is the deployment contract
// amount, decimal string, integer, decimal digit 18, decimal digit 9-18 will be ignored
// payload contract load, hex format starting from 0x and not starting from 0x
// gasPrice optional, amount, decimal string, integer, decimal digit 18, decimal digit 9-18 will be ignored; price is recommended when it is 0 or not filled from the chain
// gasLimit optional, real-time estimation for 0
func (cli *EvmClient) SendLegacyTx(key string, to *string, amount string, payload string, gasPrice string, gasLimit uint64) (string, error) {
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

	tx := cli.BuildLegacyTx(nonce, price, gasLimit, toAddr, value, data)

	signedTx, err := cli.SignLegacyTx(tx, key)
	if err != nil {
		return "", err
	}

	err = cli.SendTransaction(context.Background(), signedTx)

	return signedTx.Hash().Hex(), err
}
