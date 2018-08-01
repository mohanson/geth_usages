package main

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/mohanson/acdb"
)

const (
	cEthServer       = "https://ropsten.infura.io"
	cConfirmedNumber = 6
	cDoc             = "/tmp"
)

type Conf struct {
	BlockNumber int64
	BlockI      int64
}

func listen() error {
	rpccli, err := rpc.Dial(cEthServer)
	if err != nil {
		return err
	}
	client := ethclient.NewClient(rpccli)

	db := acdb.Doc(cDoc)
	conf := Conf{}
	db.Get("conf", &conf)

	for {
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err
		}
		if conf.BlockNumber > header.Number.Int64()-int64(cConfirmedNumber) {
			time.Sleep(time.Second * 20)
			continue
		}
		block, err := client.BlockByNumber(context.Background(), big.NewInt(conf.BlockNumber))
		if err != nil {
			return err
		}

		for i, tx := range block.Transactions() {
			if int64(i) < conf.BlockI {
				continue
			}
			// Handle Tx Tic
			log.Println(conf.BlockNumber, conf.BlockI)
			// Handle Tx Toc

			conf.BlockI = conf.BlockI + 1
			db.Set("conf", &conf)
		}

		conf.BlockNumber = conf.BlockNumber + 1
		conf.BlockI = 0
		db.Set("conf", &conf)
	}

	return nil
}

func main() {
	for {
		if err := listen(); err != nil {
			log.Println(err)
			time.Sleep(time.Second * 20)
		}
	}
}
