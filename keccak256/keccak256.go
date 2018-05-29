package main

import (
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	hash := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
	log.Println(hash.String()) // 0xa9059cbb2ab09eb219583f4a59a5d0623ade346d962bcd4e46b11da047c9049b

	hash = crypto.Keccak256Hash([]byte("name()"))
	log.Println(hash.String())
}
