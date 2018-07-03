package main

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/mohanson/simplestorage"
)

const (
	cEthServer       = "https://ropsten.infura.io"
	cTimeout         = 20
	cConfirmedNumber = 6
	cSSDir           = "/tmp"
)

func Handle(block *types.Block) error {
	log.Println(block.Number().String(), block.Hash().String())
	return nil
}

func Listen() error {
	rpccli, err := rpc.Dial(cEthServer)
	if err != nil {
		return err
	}
	client := ethclient.NewClient(rpccli)

	ss := simplestorage.New(cSSDir)
	var number int64
	if err := ss.Get("number", &number); err != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*cTimeout)
		header, err := client.HeaderByNumber(ctx, nil)
		cancel()
		if err != nil {
			return err
		}
		number = header.Number.Int64()
	}
	number = number + 1
	log.Println("From number:", number)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*cTimeout)
		header, err := client.HeaderByNumber(ctx, nil)
		cancel()
		if err != nil {
			return err
		}
		if number <= header.Number.Int64()-int64(cConfirmedNumber) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*cTimeout)
			block, err := client.BlockByNumber(ctx, big.NewInt(number))
			cancel()
			if err != nil {
				return err
			}
			if err := Handle(block); err != nil {
				log.Fatalln(err)
			}
			if err := ss.Set("number", number); err != nil {
				log.Fatalln(err)
			}
			number = number + 1
			continue
		}
		time.Sleep(time.Minute)
	}
}

func main() {
	for {
		if err := Listen(); err != nil {
			log.Println("Listen error:", err)
		}
		time.Sleep(time.Minute)
	}
}
