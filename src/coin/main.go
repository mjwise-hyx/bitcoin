package main

import (
	"core"
	"fmt"
	"strconv"
)

func main() {
	bc := core.NewBlockchain()

	bc.AddBlock("Send 1 Btc to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev.hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := core.NewProofofWork(block)
		fmt.Printf("Pow %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
