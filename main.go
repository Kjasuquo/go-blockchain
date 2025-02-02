package main

import (
	"bytes"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type Blockchain struct {
	block []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.block[len(chain.block)-1] // last block
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.block = append(chain.block, newBlock)
}

func GenesisBlock(data string) *Block {
	return CreateBlock(data, nil) // try nil later
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock("Genesis")}}
}

func main() {
	chain := InitBlockchain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.block {

		prevHash := hexutil.Encode(block.PrevHash)
		currentHash := hexutil.Encode(block.Hash)

		log.Printf("previous block hash: %v", prevHash)
		log.Printf("current block hash: %v", currentHash)
		log.Printf("block data: %v", string(block.Data))
	}
}
