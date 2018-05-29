package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

func main() {
	addressCaller := common.HexToAddress("0xe064bdf5e3e375379735a5ea4528e6099c27513f")
	caller := vm.AccountRef(addressCaller)

	addressTo := common.HexToAddress("0x0000000000000000000000000000000000000000")
	to := vm.AccountRef(addressTo)

	// caller ContractRef, object ContractRef, value *big.Int, gas uint64
	contract := vm.NewContract(caller, to, big.NewInt(0), 21000)
	log.Println(contract)

	vm.ContractRef
}
