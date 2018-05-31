package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// {"address":"0x99146Bab2bB34D9Ca49EC4f0c82De3E5789ae22e","created":824235,"decimals":0,"name":"","symbol":""}

type TokenInfo struct {
	Address  string
	Created  int64
	Decimals int64
	Name     string
	Symbol   string
}

func coderToString(s string) string {
	if strings.HasPrefix(s, "0000000000000000000000000000000000000000000000000000000000000020") {
		// string
		buf, _ := hex.DecodeString(s)
		part0 := buf[:32]
		part1 := buf[32:64]
		part2 := buf[64:]
		_ = part0

		l := binary.BigEndian.Uint64(part1[24:32])
		return string(part2[:l])
	} else if len(s) == 64 {
		// bytes32
		ctxFull, _ := hex.DecodeString(s)
		ctx := bytes.TrimRight(ctxFull, string([]byte{0}))
		return string(ctx)
	}
	return ""
}

func main() {
	f, err := os.Open("/tmp/erc20_contracts.logs")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tokenInfo := &TokenInfo{}
		if err := json.Unmarshal([]byte(scanner.Text()), tokenInfo); err != nil {
			log.Fatalln(err)
		}

		tokenInfo.Name = coderToString(tokenInfo.Name)
		tokenInfo.Symbol = coderToString(tokenInfo.Symbol)

		ctx, _ := json.Marshal(tokenInfo)
		fmt.Println(string(ctx))
	}
	if scanner.Err() != nil {
		log.Fatalln(err)
	}
}
