# 初始化环境

```sh
# 安装 geth
go install -i github.com/ethereum/go-ethereum/cmd/geth
# 同步
geth --datadir /data/geth -syncmode fast --rpc --rpcaddr=0.0.0.0 --rpcport 8545 --ws --wsport 8546 --wsaddr=0.0.0.0 --rpcapi eth,net,web3,personal,admin console 2>/tmp/geth.log
# 链接
geth --datadir /data/geth attach http://127.0.0.1:8545
```
