# 初始化环境

```sh
# 安装 geth
go install -i github.com/ethereum/go-ethereum/cmd/geth
# 同步
geth --datadir /data/geth -syncmode fast --rpc --rpcaddr=0.0.0.0 --rpcport 8545 --ws --wsport 8546 --wsaddr=0.0.0.0 --rpcapi eth,net,web3,personal,admin console 2>/tmp/geth.log
# 链接
geth --datadir /data/geth attach http://127.0.0.1:8545
```

# 二进制代码

```sh
web3.utils.keccak256('transfer(address,uint256)').slice(0, 10)
"0xa9059cbb"
```

# 判断是否是 ERC20 合约
[https://ethereum.stackexchange.com/questions/38381/how-can-i-identify-that-transaction-is-erc20-token-creation-contract](https://ethereum.stackexchange.com/questions/38381/how-can-i-identify-that-transaction-is-erc20-token-creation-contract)

```
18160ddd -> totalSupply()
70a08231 -> balanceOf(address)
dd62ed3e -> allowance(address,address)
a9059cbb -> transfer(address,uint256)
095ea7b3 -> approve(address,uint256)
23b872dd -> transferFrom(address,address,uint256)
```
