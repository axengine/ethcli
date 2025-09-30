# ethcli/v2

`ethcli/v2` 是一个用于与以太坊区块链交互的Go语言库。它封装了与 `ERC-20`、`ERC-721` 和 `ERC-1155` 代币标准相关的常用操作，并提供了交易构建、签名、发送以及其他实用工具功能，旨在简化与智能合约的交互。

## 功能特性

- **全面的代币标准支持**:
  - **ERC-20**: 标准的同质化代币功能，如查询余额、转账、授权等。
  - **ERC-721**: 非同质化代币（NFT）功能，包括查询所有者、元数据、安全转账以及各种扩展（Burnable, Pausable, Enumerable, URIStorage）。
  - **ERC-1155**: 多代币标准功能，支持批量查询余额和批量转账。
- **灵活的交易管理**:
  - 支持构建、签名和发送**传统交易 (Legacy Transactions)**。
  - 支持构建、签名和发送 **EIP-1559 动态费用交易 (Dynamic Fee Transactions)**。
- **实用工具**:
  - 生成以太坊密钥对。
  - 地址合法性校验。
  - 十六进制与字节之间的转换。

## 安装

```bash
go get github.com/axengine/ethcli/v2
```

## 使用示例

以下是一个查询 ERC-20 代币余额的简单示例：

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/axengine/ethcli/v2"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. 连接到以太坊节点
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatalf("无法连接到以太坊客户端: %v", err)
	}

	// 2. 定义代币合约地址和查询的账户地址
	tokenAddress := "0xdAC17F958D2ee523a2206206994597C13D831ec7" // USDT
	accountAddress := "0xYourAccountAddress"

	// 3. 调用方法查询余额
	balance, err := ethcli.ERC20BalanceOf(context.Background(), client, tokenAddress, accountAddress, nil)
	if err != nil {
		log.Fatalf("无法获取代币余额: %v", err)
	}

	// 4. 获取代币精度以便正确显示
	decimals, err := ethcli.ERC20Decimals(context.Background(), client, tokenAddress, nil)
	if err != nil {
		log.Fatalf("无法获取代币精度: %v", err)
	}

	// 格式化余额
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	value := new(big.Float).Quo(fBalance, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)))

	fmt.Printf("账户 %s 的余额为: %f USDT
", accountAddress, value)
}
```

## API 方法概览

### ERC-20

- `ERC20Name(token string) (string, error)`: 获取代币名称。
- `ERC20Symbol(token string) (string, error)`: 获取代币符号。
- `ERC20Decimals(token string) (uint8, error)`: 获取代币精度。
- `ERC20TotalSupply(token string) (*big.Int, error)`: 获取代币总供应量。
- `ERC20BalanceOf(token, owner string) (*big.Int, error)`: 查询账户余额。
- `ERC20Allowance(token, owner, spender string) (*big.Int, error)`: 查询授权额度。
- `ERC20Transfer(key, token, to, value string) (string, error)`: 发送代币。
- `ERC20Approve(key, token, spender, value string) (string, error)`: 授权给第三方。
- `ERC20TransferFrom(key, token, from, to, value string) (string, error)`: 从指定账户转出代币。
- `ERC20Mint(key, token, to, value string) (string, error)`: 铸造新币。
- `ERC20Burn(key, token, value string) (string, error)`: 销毁代币。
- `ERC20BurnFrom(key, token, owner, value string) (string, error)`: 从指定账户销毁代币。

### ERC-721

- `ERC721BalanceOf(token, owner string) (*big.Int, error)`: 查询账户拥有的NFT数量。
- `ERC721OwnerOf(token string, tokenId *big.Int) (string, error)`: 查询NFT的所有者。
- `ERC721GetApproved(token string, tokenId *big.Int) (string, error)`: 获取指定NFT的授权地址。
- `ERC721IsApprovedForAll(token, owner, operator string) (bool, error)`: 检查是否已授权所有NFT给操作者。
- `ERC721TokenURI(token string, tokenId *big.Int) (string, error)`: 获取NFT的元数据URI。
- `ERC721TransferFrom(key, token, from, to string, tokenId *big.Int) (string, error)`: 转移NFT。
- `ERC721Approve(key, token, to string, tokenId *big.Int) (string, error)`: 授权单个NFT。
- `ERC721SetApprovalForAll(key, token, operator string, approved bool) (string, error)`: 授权或取消所有NFT的操作权限。
- `ERC721Mint(...)`: 提供多种铸造NFT的方法。
- `ERC721Burn(key, token string, tokenId *big.Int) (string, error)`: 销毁NFT。
- `ERC721Pause(key, token string) (string, error)`: 暂停合约。
- `ERC721Unpause(key, token string) (string, error)`: 取消暂停。

### ERC-1155

- `ERC1155BalanceOf(token, owner string, tokenId *big.Int) (*big.Int, error)`: 查询单个代币ID的余额。
- `ERC1155BalanceOfBatch(token string, owners []string, tokenIds []*big.Int) ([]*big.Int, error)`: 批量查询余额。
- `ERC1155IsApprovedForAll(token, owner, operator string) (bool, error)`: 检查是否已授权。
- `ERC1155SafeTransferFrom(...)`: 安全转移代币。
- `ERC1155SafeBatchTransferFrom(...)`: 批量安全转移代币。
- `ERC1155Mint(...)`: 铸造代币。
- `ERC1155Burn(...)`: 销毁代币。

### 交易与工具

- `SendLegacyTx(...)`: 发送传统交易。
- `SendDynamicFeeTx(...)`: 发送EIP-1559交易。
- `GenKey() (string, string, string, error)`: 生成新的以太坊账户密钥。
- `ValidAddress(address string) bool`: 验证地址格式是否正确。
