package core

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	Bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Printf("Usage:")
	fmt.Printf("	addblock -data BLOCK_DaTA - add a block to the blockchain \n")
	fmt.Printf("	printchain - print all the blocks of the blockchain \n")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.Bc.AddBlock(data)
	fmt.Println("Success~!")
}

func (cli *CLI) printChain() {
	bci := cli.Bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("{ \n")
		fmt.Printf("	PrevBlockHash ; %x \n", block.PrevBlockHash)
		fmt.Printf("	Data ; %s \n", block.Data)
		fmt.Printf("	Hash ; %x \n", block.Hash)
		fmt.Printf("	Timestamp ; %v \n", block.Timestamp)

		pow := NewProofOfWork(block)
		fmt.Printf("	PoW: %s \n", strconv.FormatBool(pow.Validate()))

		fmt.Printf("{ \n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// 解析命令行参数,并执行命令
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) addBlock33(data string) {

}
