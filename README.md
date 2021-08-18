# ethcli
- 包装go-ethereum ethclient.Client，提供内置的web3 JSON RPC方法
- 封装离线签名交易接口（`BuildTx`、`SignTx`、`SendTx`），封装高级别的`SendMondoTx`接口
- 封装ORC20/ERC20接口
- 封装ORC721/ERC721接口，封装自定义ORC721/ERC721 `mint`接口
- 提供地址生成`GenKey`、校验`ValidAddress`、资产精度转换方法`ToWei`,`ToEther`，十六进制数据转换`HexToBytes` `BytesToHex` `BytesToHexWith0x`方法

## ORC20
- 实现标准ERC20接口

## ORC721
- 实现`openzeppelin-contract/contracts/token/ERC721` 和 `openzeppelin-contract/contracts/token/ERC721/extensions` 相关接口
- 实现自定义`uint256 mint(address to)` 和 `uint256 mint(address to,string memory _tokenURI)` 接口
- `string tokenURI(uint256 tokenId)`方法由`ERC721URIStorage`和`ERC721Metadata`共用，mint返回的tokenId只能通过hash查询交易结果，从交易事件中获取