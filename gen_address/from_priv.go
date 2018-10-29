package main

import (
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privKeyStr := "3a4323bb01c89f7bb800fa59b3a8ed81dab540110242f21d914e912371da78f9"
	privKeyBuf, _ := hex.DecodeString(privKeyStr)
	privKey, _ := crypto.ToECDSA(privKeyBuf)
	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	log.Printf("%s %s", privKeyStr, addr.Hex())
}
