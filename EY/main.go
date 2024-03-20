package main

import (
	"fmt"
)

func main() {

	coinConfig := NewCoinConfig()

	// 创建新的区块链实例
	blockchain := NewBlockchain(coinConfig)

	// 向区块链中添加一些区块
	blockchain.AddBlock("First Block")
	blockchain.AddBlock("Second Block")
	blockchain.AddBlock("Third Block")

	// 打印区块链中的所有区块信息
	for _, block := range blockchain.blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}

}
