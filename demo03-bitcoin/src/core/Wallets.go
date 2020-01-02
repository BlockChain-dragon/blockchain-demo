package core

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletsFile = "wallets.dat" //存储钱包数据的本地文件名

//定义一个钱包的集合，存储多个钱包对象
type Wallets struct {
	WalletMap map[string]*Wallet
}

//创建钱包集合
func (ws *Wallets) CreateNewWallets() {
	wallet := NewWallet()
	var address []byte
	address = wallet.GetAddress()
	//创建想要的比特币地址
	//for {
	//	address = wallet.GetAddress()
	//	fmt.Printf("\r%s", address)
	//
	//	if strings.Contains(string(address), "bru") {
	//		break
	//	}
	//}
	fmt.Printf("创建的钱包地址：%s\n", address)
	ws.WalletMap[string(address)] = wallet
	//将钱包集合存入到本地文件中
	ws.SaveFile()
}

//将钱包对象，存入到本地文件中
func (ws *Wallets) SaveFile() {
	//1.将ws对象的数据--->byte[]
	var buf bytes.Buffer
	//序列化的过程中：被序列化的对象 中包含了接口，那么将口需要注册
	gob.Register(elliptic.P256()) //Curve
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}

	wsBytes := buf.Bytes()

	//2.将数据存储到文件中，
	//注意：该方法的实现：ioutil.WriteFile，覆盖写数据
	err = ioutil.WriteFile(walletsFile, wsBytes, 0644)
	if err != nil {
		log.Panic(err)
	}
}

//提供一个函数，用于获取一个钱包的集合
/*
思路：
	读取本地的钱包文件，如果文件存在，直接获取
	如果文件不存在，创建一个空的钱包对象
*/
func GetWallets() *Wallets {
	//step1：钱包文件不存在
	if _, err := os.Stat(walletsFile); os.IsNotExist(err) {
		fmt.Println("区块链钱包不存在")
		//创建钱包集合
		wallets := &Wallets{}
		wallets.WalletMap = make(map[string]*Wallet)
		return wallets
	}

	//step2:钱包文件存在
	//读取本地的钱包文件中的数据---->反序列得到钱包集合对象
	wsBytes, err := ioutil.ReadFile(walletsFile)
	if err != nil {
		log.Panic(err)
	}

	gob.Register(elliptic.P256()) //Curve

	//将数据，变成钱包集合对象
	var wallets Wallets
	reader := bytes.NewReader(wsBytes)
	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}
	return &wallets

}
