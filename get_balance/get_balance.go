package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	rpcDial, err := rpc.Dial("http://10.0.5.198:8547")
	if err != nil {
		log.Fatalln(err)
	}
	client := ethclient.NewClient(rpcDial)

	r, err := client.BalanceAt(context.Background(), common.HexToAddress("0x3aa9e3ab43f7149f3842d67863ecd80adba3447e"), nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r)
}
