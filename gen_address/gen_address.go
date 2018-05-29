package main

import (
	"crypto/rand"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	buf := make([]byte, 32)
	for i := 0; i < 100; i++ {
		rand.Read(buf)
		key, err := crypto.ToECDSA(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		addr := crypto.PubkeyToAddress(key.PublicKey)
		log.Printf("0x%x %s", buf, addr.Hex())
	}
}
