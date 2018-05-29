package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	set := map[string]struct{}{}

	f, err := os.Open("eth_address_uniq.log")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		set[line] = struct{}{}
	}
	f.Close()
	log.Println("Count of addresses", len(set))

	logFile, err := os.OpenFile("eth_priv.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer logFile.Close()
	w := io.MultiWriter(os.Stdout, logFile)

	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, 32)
	for {
		rand.Read(buf)
		key, err := crypto.HexToECDSA(hex.EncodeToString(buf))
		if err != nil {
			log.Fatalln(err)
		}
		addr := crypto.PubkeyToAddress(key.PublicKey)
		_, ok := set[addr.String()]
		if ok {
			w.Write([]byte(fmt.Sprintf("0x%x %s\n", buf, addr.String())))
		}
	}
}
