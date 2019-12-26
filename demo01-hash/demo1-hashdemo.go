package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math"
)

func main() {
	var max222 int64 = math.MaxInt64
	// 9223372036854775807
	fmt.Print(max222)

	//calculateHash("test`1")
}

func calculateHash(toBeHashed string) string {
	hashInBytes := sha256.Sum256([]byte(toBeHashed))
	hashInStr := hex.EncodeToString(hashInBytes[:])
	log.Printf("%s , %s", toBeHashed, hashInStr)
	return hashInStr
}
