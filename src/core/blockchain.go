package core

//BlockChain keeps a sequence of Blocks
type Blockchain struct {
	Blocks []*Block
}

// AddBlock saves provided datea as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	//新的区块保存前面区块的信息
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
