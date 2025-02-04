package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go-blockchain/blockchain"
	"log"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.Block {

		prevHash := hexutil.Encode(block.PrevHash)
		currentHash := hexutil.Encode(block.Hash)

		log.Printf("previous block hash: %v", prevHash)
		log.Printf("current block hash: %v", currentHash)
		log.Printf("block data: %v", string(block.Data))
		log.Printf("block nonce: %v", block.Nonce)

		//log.Printf("previous block hash: %v", string(block.PrevHash))
		//log.Printf("current block hash: %v", string(block.Hash))
		//log.Printf("block data: %v", string(block.Data))
		//log.Printf("block nonce: %v", block.Nonce)

		pow := blockchain.NewProofOfWork(block)
		log.Printf("pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
