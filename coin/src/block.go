package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"time"
)

// Block 定义区块链的区块结构
type Block struct {
	Index         int64
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int64
	Reward        int64
}

// NewBlock 创建新的区块
func NewBlock(transactions []*Transaction, prevBlock *Block, reward int64) *Block {
	block := &Block{
		Index:         prevBlock.Index + 1,
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlock.Hash, // 使用 prevBlock.Hash 而不是 prevBlockHash（如果 prevBlock 不是 nil）
		Reward:        reward,
		Nonce:         0,
		Hash:          []byte{},
	}

	// 计算区块哈希值
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

// Serialize 序列化区块为JSON格式
func (b *Block) Serialize() []byte {
	jsonData, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonData
}

// NewGenesisBlock 创建创世区块
func NewGenesisBlock(transactions []*Transaction, reward int64) *Block {
	return NewBlock(transactions, nil, reward)
}

// ProofOfWork Proof of Work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork 初始化Proof of Work
func NewProofOfWork(b *Block) *ProofOfWork {
	// 这里假设target是一个固定值，实际中应该根据网络难度动态生成
	target := big.NewInt(0)
	target.Lsh(target, 256-24) // 设置target为2^240，简化难度
	return &ProofOfWork{b, target}
}

// Run 执行Proof of Work，寻找合适的nonce值使得区块哈希值满足难度要求
func (pow *ProofOfWork) Run() (int64, []byte) {
	var hash [32]byte
	var hashInt big.Int

	nonce := int64(0)
	for {
		data := bytes.Join(
			[][]byte{
				pow.block.PrevBlockHash,
				pow.block.Serialize(),
				[]byte(fmt.Sprintf("%d", nonce)),
			},
			[]byte{},
		)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		// 如果当前哈希小于目标值，则找到了合适的nonce
		if hashInt.Cmp(pow.target) == -1 {
			break
		}
		nonce++

		// 添加超时或最大迭代次数以防止无限循环
		// 这里简化为仅检查nonce是否超过一个合理值
		if nonce > math.MaxInt64/2 {
			fmt.Println("Proof of Work exceeded maximum nonce without success")
			return nonce, hash[:]
		}
	}

	return nonce, hash[:]
}

// ValidateBlock 验证区块的合法性
func ValidateBlock(block, prevBlock *Block) bool {
	if prevBlock != nil && !bytes.Equal(block.PrevBlockHash, prevBlock.Hash) {
		return false
	}
	pow := NewProofOfWork(block)
	_, targetHash := pow.Run()

	return bytes.Equal(block.Hash, targetHash)
}

func main() {
	fmt.Printf("%v\n", ValidateBlock)
	fmt.Printf("%v\n", NewGenesisBlock)
}
