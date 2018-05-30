package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/number"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	nodeAddress = "http://10.0.5.198:8547"
	tic         = 800000
	toc         = 5600000
	// tic         = 5622390
	// toc         = 5622399
	// tic        = 3904410
	// toc        = 3904420
	saveToFile = "erc20_contracts.logs"
)

func getContractName(client *ethclient.Client, address *common.Address) ([]byte, error) {
	funcHex := crypto.Keccak256([]byte("name()"))
	funcHexPrefix := funcHex[:4]

	cnt, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   address,
		Data: funcHexPrefix,
	}, nil)
	if err != nil {
		return []byte{}, err
	}
	return cnt, nil
}

func getContractSymbol(client *ethclient.Client, address *common.Address) ([]byte, error) {
	funcHex := crypto.Keccak256([]byte("symbol()"))
	funcHexPrefix := funcHex[:4]

	cnt, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   address,
		Data: funcHexPrefix,
	}, nil)
	if err != nil {
		return []byte{}, err
	}
	return cnt, nil
}

func getContractDecimal(client *ethclient.Client, address *common.Address) (*number.Number, error) {
	funcHex := crypto.Keccak256([]byte("decimals()"))
	funcHexPrefix := funcHex[:4]

	cnt, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   address,
		Data: funcHexPrefix,
	}, nil)
	if err != nil {
		return number.Uint256(0), err
	}

	d := number.Uint256(0)
	d.SetBytes(cnt)
	return d, nil
}

func main() {
	defer log.Println("done")
	rpcDial, err := rpc.Dial(nodeAddress)
	if err != nil {
		log.Fatalln(err)
	}
	client := ethclient.NewClient(rpcDial)

	file, err := os.OpenFile(saveToFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	w := io.MultiWriter(os.Stdout, file)

	for i := tic; i < toc; i++ {
		if i%10000 == 0 {
			log.Println("Current Block is", i)
		}

		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			log.Println("Current Block is", i)
			log.Fatalln(err)
		}

		for _, tx := range block.Transactions() {
			// it's a contract creation
			if tx.To() == nil {
				// get contract address
				receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					log.Println("Current Block is", i)
					log.Fatalln(err)
				}

				// get contract code
				jsonTx, err := tx.MarshalJSON()
				if err != nil {
					log.Fatalln(err)
				}
				obj := map[string]string{}
				if err := json.Unmarshal(jsonTx, &obj); err != nil {
					log.Println("Current Block is", i)
					log.Fatalln(err)
				}
				code := obj["input"]

				// is erc20 token?
				// https://ethereum.stackexchange.com/questions/38381/how-can-i-identify-that-transaction-is-erc20-token-creation-contract
				if strings.Contains(code, "18160ddd") &&
					strings.Contains(code, "70a08231") &&
					strings.Contains(code, "dd62ed3e") &&
					strings.Contains(code, "a9059cbb") &&
					strings.Contains(code, "095ea7b3") &&
					strings.Contains(code, "23b872dd") {

					name, err := getContractName(client, &receipt.ContractAddress)
					if err != nil {
						log.Println("Current Block is", i)
						log.Fatalln(err)
					}

					symbol, err := getContractSymbol(client, &receipt.ContractAddress)
					if err != nil {
						log.Println("Current Block is", i)
						log.Fatalln(err)
					}

					decimals, err := getContractDecimal(client, &receipt.ContractAddress)
					if err != nil {
						log.Println("Current Block is", i)
						log.Fatalln(err)
					}

					data := map[string]interface{}{}
					data["address"] = receipt.ContractAddress.String()
					data["name"] = hex.EncodeToString(name)
					data["symbol"] = hex.EncodeToString(symbol)
					data["decimals"], _ = strconv.Atoi(decimals.String())
					data["created"] = i

					line, _ := json.Marshal(data)
					w.Write(line)
					w.Write([]byte{'\n'})
				}
			}
		}
	}
}
