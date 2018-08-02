package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	cEthServer = "https://mainnet.infura.io"
)

func listen() error {
	rpccli, err := rpc.Dial(cEthServer)
	if err != nil {
		return err
	}
	client := ethclient.NewClient(rpccli)

	block, err := client.BlockByNumber(context.Background(), big.NewInt(6073453))
	if err != nil {
		return err
	}
	for _, tx := range block.Transactions() {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Println(err)
		}

		if len(receipt.Logs) == 0 {
			continue
		}

		for _, txlog := range receipt.Logs {
			if len(txlog.Topics) == 0 {
				continue
			}
			if txlog.Topics[0].String() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
				fromAddr := "0x" + txlog.Topics[1].String()[26:66]
				toAddr := "0x" + txlog.Topics[2].String()[26:66]
				value := big.NewInt(0)
				value.SetBytes(txlog.Data)
				log.Printf("发生了 ERC20 交易. 合约地址=%v TxHash=%v from=%v to=%v value=%v", txlog.Address.String(), tx.Hash().String(), fromAddr, toAddr, value)
			}
		}
	}
	return nil
}

func main() {
	if err := listen(); err != nil {
		log.Fatalln(err)
	}
}
