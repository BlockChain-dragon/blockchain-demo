package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func main() {
	calculateHash("test`1")
}

func calculateHash(toBeHashed string) string {
	hashInBytes := sha256.Sum256([]byte(toBeHashed))
	hashInStr := hex.EncodeToString(hashInBytes[:])
	log.Printf("%s , %s", toBeHashed, hashInStr)
	return hashInStr
}
