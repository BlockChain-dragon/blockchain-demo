package main

import (
	"./core"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {

	//level2
	//level2()

	// leve 3 pow原理
	//proof()

	//leve 4
	cli := core.CLI{}
	cli.Run()
	//arr := "[HtJve4MwW4LVko3yjkj3UUvzXsuGHb1Yq]"
	//srtArr := core.JSONToArray(arr)
	//for _,value := range srtArr  {
	//	println(value)
	//}

}

const targetBits = 24

func proof() {
	data1 := []byte("i like donuts")
	data2 := []byte("i like donutscasd")
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	fmt.Printf("%x \n", sha256.Sum256(data1))
	fmt.Printf("%064x \n", target)
	fmt.Printf("%x \n", sha256.Sum256(data2))

}
