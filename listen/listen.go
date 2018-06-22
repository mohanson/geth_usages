package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// rpcDial, err := rpc.Dial("https://mainnet.infura.io/abAWsazRr6zO8zmW8J4i")
	rpcDial, err := rpc.Dial("wss://mainnet.infura.io/ws")
	if err != nil {
		log.Fatalln(err)
	}
	client := ethclient.NewClient(rpcDial)

	ch := make(chan *types.Header, 4)
	sub, err := client.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case hea := <-ch:
			log.Println(hea.Number.String(), hea.Hash().String())
		case err := <-sub.Err():
			log.Fatalln(err)
		}
	}
}
