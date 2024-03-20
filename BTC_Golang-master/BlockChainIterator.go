package main

import (
	"github.com/boltdb/bolt"
	"log"
)

/**
区块链迭代器
*/

// BlockChainIterator 区块链迭代器的结构
type BlockChainIterator struct {
	db           *bolt.DB
	currentPoint []byte //指向当前的区块
}

// NewBlockChainIterator 迭代器创建函数
func NewBlockChainIterator(bc *BlockChain) BlockChainIterator {
	return BlockChainIterator{
		db:           bc.db,
		currentPoint: bc.tail,
	}
}

// GetBlockAndMoveLeft 迭代器访问函数
func (it *BlockChainIterator) GetBlockAndMoveLeft() *Block {
	var block *Block

	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBecket))

		if bucket == nil {
			log.Panic("bucket should not be nil !")
		} else {
			//根据当前的current_pointer获取block
			data := bucket.Get(it.currentPoint)
			block = Deserialize(data)
			it.currentPoint = block.PreHash //将游标（指针）左移
		}

		return nil
	})

	return block
}
