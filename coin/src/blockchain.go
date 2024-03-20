package main

import (
	"encoding/json"
	"io/ioutil"
)
 
// Blockchain 结构体
type Blockchain struct {
	blocks []*Block
}

// NewBlockchain 初始化一个新的区块链
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{}}
}

// AddBlock 将新的区块添加到区块链中
func (bc *Blockchain) AddBlock(newBlock *Block) {
	bc.blocks = append(bc.blocks, newBlock)
}

// SaveBlockchain 将区块链保存到文件中
func (bc *Blockchain) SaveBlockchain(filename string) error {
	jsonData, err := json.Marshal(bc.blocks)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadBlockchain 从文件中加载区块链
func LoadBlockchain(filename string) (*Blockchain, error) {
	var blocks []*Block
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileContent, &blocks)
	if err != nil {
		return nil, err
	}
	return &Blockchain{blocks: blocks}, nil
}
