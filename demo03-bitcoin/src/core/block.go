package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

/*

时间戳
数据
前hash
hash
*/
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

//set hash
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}

//创世
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
