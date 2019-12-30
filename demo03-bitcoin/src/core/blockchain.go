package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// level 1
//type Blockchain struct {
//	Blocks []*Block
//}
//level 2
type Blockchain struct {
	tip []byte
	Db  *bolt.DB
}

type BlockchainIterator struct {
	CurrentHash []byte
	Db          *bolt.DB
}

// level 1
//add block
//func (bc *Blockchain) AddBlock(data string) {
//	prevBlock := bc.Blocks[len(bc.Blocks)-1]
//	newBlock := NewBlock(data, prevBlock.Hash)
//	bc.Blocks = append(bc.Blocks, newBlock)
//}
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put([]byte(newBlock.Hash), newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash

		return nil
	})

}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.Db}
	return bci
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodeBlock := b.Get([]byte(i.CurrentHash))
		block = DeserializeBlock(encodeBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	i.CurrentHash = block.PrevBlockHash

	return block
}

// level 1
//add block
//func (bc *Blockchain) PrintBlockchain() {
//	for _, block := range bc.Blocks {
//		fmt.Printf("{ \n")
//		fmt.Printf("	PrevBlockHash ; %x \n", block.PrevBlockHash)
//		fmt.Printf("	Data ; %s \n", block.Data)
//		fmt.Printf("	Hash ; %x \n", block.Hash)
//		fmt.Printf("	Timestamp ; %v \n", block.Timestamp)
//
//		pow := NewProofOfWork(block)
//		fmt.Printf("PoW: %s \n", strconv.FormatBool(pow.Validate()))
//
//		fmt.Printf("{ \n")
//	}
//}

// level 1 创建新区块-内存
//func NewBlockchain() *Blockchain {
//	return &Blockchain{[]*Block{NewGenesisBlock()}}
//}
//level 2  创建新区块-文件
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil { // 读取不到文件
			fmt.Println("No existing blockchain found . Creating a new one...")
			genesis := NewGenesisBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			// l 表示 leader
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = genesis.Hash
		} else { //读取到文件
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}
