package core

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("	createblockchain -address ADDRESS   -- Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("	createwallet                        -- Generates a new key-pair and saves it into the wallet file")
	fmt.Println("	getbalance -address ADDRESS         -- Get balance of ADDRESS")
	fmt.Println("	listaddresses                       -- Lists all addresses from the wallet file")
	fmt.Println("	printchain                          --print all the blocks of the blockchain ")
	fmt.Println("	send -from [FROMADDRESSES] -to [TOADDRESSES] -amount [AMOUNT] -- Send AMOUNT of coins from FROMADDRESS to TOADDRESS ")
}

//判断终端输入的参数的长度
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// 解析命令行参数,并执行命令
func (cli *CLI) Run() {
	cli.validateArgs()

	//1.创建flagset命令对象
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	testMethodCmd := flag.NewFlagSet("test", flag.ExitOnError)

	//2.设置命令后的参数对象
	getBalanceAddress := getBalanceCmd.String("address", "", "The address get balance ")
	createBlockchainAddress := createBlockchainCmd.String("address", "GenesisBlock", "The address to creat blockchain ")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.String("amount", "", "Amount to send")

	//3.解析命令对象
	switch os.Args[1] {
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	// level 4
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "test":
		err := testMethodCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "help":
		cli.printUsage()
		os.Exit(1)
	default:
		cli.printUsage()
		os.Exit(1)
	}
	//4.根据终端输入的命令执行对应的功能
	//4.1 创建钱包--->交易地址
	if createWalletCmd.Parsed() {
		cli.CreateWallet()
	}
	//4.2 获取钱包地址
	if listAddressesCmd.Parsed() {
		cli.ListAddresses()
	}

	//4.3 创建创世区块
	if createBlockchainCmd.Parsed() {
		if !IsValidAddress([]byte(*createBlockchainAddress)) {
			fmt.Println("地址无效，无法创建创世前区块")
			cli.printUsage()
			os.Exit(1)
		}
		cli.CreateBlockChain(*createBlockchainAddress)
	}
	//4.4 转账交易
	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount == "" {
			fmt.Println("转账信息有误")
			cli.printUsage()
			os.Exit(1)
		}
		//添加区块
		from := JSONToArray(*sendFrom)     //[]string
		to := JSONToArray(*sendTo)         //[]string
		amount := JSONToArray(*sendAmount) //[]string
		for i := 0; i < len(from); i++ {
			if !IsValidAddress([]byte(from[i])) || !IsValidAddress([]byte(to[i])) {
				fmt.Println("地址无效，无法转账")
				cli.printUsage()
				os.Exit(1)
			}
		}

		cli.Send(from, to, amount)
	}
	//4.5 查询余额
	if getBalanceCmd.Parsed() {
		if !IsValidAddress([]byte(*getBalanceAddress)) {
			fmt.Println("查询地址有误")
			cli.printUsage()
			os.Exit(1)
		}
		cli.GetBalance(*getBalanceAddress)
	}

	//4.6 打印区块信息
	if printChainCmd.Parsed() {
		cli.PrintChains()
	}

	if testMethodCmd.Parsed() {
		cli.TestMethod()
	}

}
