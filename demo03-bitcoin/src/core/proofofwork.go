package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce int64 = math.MaxInt32
)

const targetBits = 20

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{pow.block.PrevBlockHash, pow.block.Data, IntToHex(int64(targetBits)), IntToHex(int64(nonce))},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	var nonce int64

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
		data := pow.prepareData(int(nonce))

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return int(nonce), hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

//共识
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}
