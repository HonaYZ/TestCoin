package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block 结构体表示区块链中的一个区块
type Block struct {
	Index     int
	Timestamp string
	Data      string // 区块中存储的数据，可以是交易信息或其他内容
	PrevHash  string // 前一个区块的哈希值
	Hash      string // 当前区块的哈希值
}

// Blockchain 结构体表示整个区块链
type Blockchain struct {
	blocks []*Block
	config *CoinConfig
}

// NewBlockchain 创建一个新的区块链实例
func NewBlockchain(config *CoinConfig) *Blockchain {
	genesisBlock := createGenesisBlock(config)
	return &Blockchain{blocks: []*Block{genesisBlock}, config: config}
}

// createGenesisBlock 创建一个创世区块
func createGenesisBlock(config *CoinConfig) *Block {
	return createBlock(0, "Genesis Block", "", config)
}

// createBlock 创建一个新的区块
func createBlock(index int, data string, prevHash string, config *CoinConfig) *Block {
	timestamp := time.Now().Format(time.RFC3339)
	hash := calculateHash(index, timestamp, data, prevHash, config)

	return &Block{
		Index:     index,
		Timestamp: timestamp,
		Data:      data,
		PrevHash:  prevHash,
		Hash:      hash,
	}
}

// calculateHash 计算区块的哈希值
func calculateHash(index int, timestamp string, data string, prevHash string, config *CoinConfig) string {
	record := strconv.Itoa(index) + timestamp + data + prevHash + config.Symbol
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// AddBlock 向区块链中添加一个新的区块
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := createBlock(prevBlock.Index+1, data, prevBlock.Hash, bc.config)
	bc.blocks = append(bc.blocks, newBlock)
}

// ... 其他区块链相关的方法 ...
