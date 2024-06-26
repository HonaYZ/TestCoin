package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"github.com/base58"
	"io/ioutil"
	"log"
	"os"
)

/**
保存所有的wallet以及它的地址
*/

type Wallets struct {
	WalletsMap map[string]*Wallet //map[地址]钱包
}

// NewWallets 创建方法，返回当前所有钱包的实例
func NewWallets() *Wallets {
	var wallets Wallets
	wallets.WalletsMap = make(map[string]*Wallet)
	wallets.loadFile()
	return &wallets
}

// CreateWallet 创建钱包
func (wallets *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()

	wallets.WalletsMap[address] = wallet
	wallets.saveToFile()
	return address
}

// 保存方法，把新建的wallet添加进去
func (wallets *Wallets) saveToFile() {
	var buffer bytes.Buffer

	//panic: gob: type not registered for interface: elliptic.p256Curve
	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(wallets)
	if err != nil {
		log.Panic(err)
	}

	ioutil.WriteFile(walletsFile, buffer.Bytes(), 0600)
}

// 读取文件方法，把所有的wallet读出来
func (wallets *Wallets) loadFile() {
	//在读取之前，要先确认文件是否在，如果不存在，直接退出
	_, err := os.Stat(walletsFile)
	if os.IsNotExist(err) {
		return
	}

	//读取内容
	content, err := ioutil.ReadFile(walletsFile)
	if err != nil {
		log.Panic(err)
	}

	//解码
	//panic: gob: type not registered for interface: elliptic.p256Curve
	gob.Register(elliptic.P256())

	decoder := gob.NewDecoder(bytes.NewReader(content))

	var wsLocal Wallets

	err = decoder.Decode(&wsLocal)
	if err != nil {
		log.Panic(err)
	}

	//wallets = &wsLocal
	//对于结构来说，里面有map的，要指定赋值，不要再最外层直接赋值
	wallets.WalletsMap = wsLocal.WalletsMap
}

// ListAddresses 遍历钱包
func (wallets *Wallets) ListAddresses() []string {
	var addresses []string
	//遍历钱包，将所有的key取出来返回
	for address := range wallets.WalletsMap {
		addresses = append(addresses, address)
	}

	return addresses
}

// GetPubKeyFromAddress 通过地址返回公钥的哈希值
func GetPubKeyFromAddress(address string) []byte {
	//1. 解码
	addressByte := base58.Decode(address) //25字节
	len := len(addressByte)

	//2. 截取出公钥哈希：去除version（1字节），去除校验码（4字节）
	pubKeyHash := addressByte[1 : len-4]

	return pubKeyHash
}
