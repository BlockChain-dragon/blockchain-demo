package main

import (
	"../core"
)

func main() {
	bc := core.NewBlockchain()
	bc.AddBlock("hello world ~!")
	bc.AddBlock("hello flower")

	bc.PrintBlockchain()

}
