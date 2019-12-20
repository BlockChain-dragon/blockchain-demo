package main

import "../core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("hello world")
	bc.SendData("hello www")
	bc.Print()

}
