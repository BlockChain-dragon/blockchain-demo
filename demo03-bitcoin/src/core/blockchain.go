package core

import (
	"fmt"
	"strconv"
)

type Blockchain struct {
	Blocks []*Block
}

//add block
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

//add block
func (bc *Blockchain) PrintBlockchain() {
	for _, block := range bc.Blocks {
		fmt.Printf("{ \n")
		fmt.Printf("	PrevBlockHash ; %x \n", block.PrevBlockHash)
		fmt.Printf("	Data ; %s \n", block.Data)
		fmt.Printf("	Hash ; %x \n", block.Hash)
		fmt.Printf("	Timestamp ; %v \n", block.Timestamp)

		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s \n", strconv.FormatBool(pow.Validate()))

		fmt.Printf("{ \n")
	}
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
