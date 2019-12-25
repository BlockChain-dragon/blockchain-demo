package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  // 区块编号
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上一个区块hash
	Hash          string //当前区块哈希值
	Data          string //区块数据
}

// 计算区块hash
func CalculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInByte := sha256.Sum256([]byte(blockData))
	hashInStr := hex.EncodeToString(hashInByte[:])
	return hashInStr
}

// 生成新区块
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock
}

//创世区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}
