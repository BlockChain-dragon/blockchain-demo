package core

type Blockchain struct {
	Blocks []*Block
}

//add block
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

//创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
