package main

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	nb10e9  = new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)
	nb10e18 = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
)

var (
	// You should enter your private key
	cFromPrivKey = "----------------------------------------------------------------"
	cFromAddress = "0xeb1379888f6117386043b1e50aafa983006958d8"
	cToAddress   = "0xe064bdF5E3E375379735A5EA4528E6099c27513f"
	cAmount      = 10
	cGasPrice    = 80000
	cGasLimit    = 80000
)

func main() {
	rpcDial, err := rpc.Dial("https://ropsten.infura.io/abAWsazRr6zO8zmW8J4i")
	if err != nil {
		log.Fatalln(err)
	}
	client := ethclient.NewClient(rpcDial)

	fromAddress := common.HexToAddress(cFromAddress)
	toAddress := common.HexToAddress(cToAddress)

	buf, err := hex.DecodeString(cFromPrivKey)
	if err != nil {
		log.Fatalln(err)
	}
	fromPrivKey, err := crypto.ToECDSA(buf)
	if err != nil {
		log.Fatalln(err)
	}

	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Current nonce is", nonce)

	tx := types.NewTransaction(
		nonce,
		toAddress,
		new(big.Int).Mul(big.NewInt(int64(cAmount)), nb10e18),
		uint64(cGasLimit),
		new(big.Int).Mul(big.NewInt(int64(cGasPrice)), nb10e9),
		[]byte{},
	)

	signer := types.HomesteadSigner{}
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), fromPrivKey)
	if err != nil {
		log.Fatalln(err)
	}
	txSigned, err := tx.WithSignature(signer, signature)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(txSigned.Hash().String())
	if err := client.SendTransaction(context.Background(), txSigned); err != nil {
		log.Fatalln(err)
	}
}
