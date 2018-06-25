package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	rpcDial, err := rpc.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalln(err)
	}
	client := ethclient.NewClient(rpcDial)

	r, err := client.BalanceAt(context.Background(), common.HexToAddress("0x0000000000000000000000000000000000000000"), nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r)
}
