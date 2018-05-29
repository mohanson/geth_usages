package main

import (
	"context"
	"io"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	host = "http://127.0.0.1:8547"
	i0   = 0
	i1   = 5600000
)

func main() {
	log.Println("Listen on", host)
	rpcDial, err := rpc.Dial(host)
	if err != nil {
		log.Fatalln(err)
	}
	client := ethclient.NewClient(rpcDial)

	logFile, err := os.OpenFile("eth_address.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer logFile.Close()
	w := io.MultiWriter(os.Stdout, logFile)

	for i := i0; i < i1; i++ {
		log.Println(i)
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		block, err := client.BlockByNumber(ctx, big.NewInt(int64(i)))
		cancel()
		if err != nil {
			continue
		}

		transactions := block.Transactions()
		for _, transaction := range transactions {
			address := transaction.To()
			if address != nil {
				w.Write([]byte(address.String() + "\n"))
			}
		}
	}
}
