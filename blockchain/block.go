package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

type Blockchain struct {
	Block []*Block
}

// it is already derived inside proof of work

//func (b *Block) DeriveHash() {
//	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
//	hash := sha256.Sum256(info)
//	b.Hash = hash[:]
//}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	//block.DeriveHash()
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Block[len(chain.Block)-1] // last block
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Block = append(chain.Block, newBlock)
}

func GenesisBlock(data string) *Block {
	return CreateBlock(data, []byte{}) // try nil later
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock("Genesis")}}
}
