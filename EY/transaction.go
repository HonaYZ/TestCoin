package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

// Transaction 结构体表示一个交易
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
	Time int64
}

// TXInput 结构体表示交易的输入
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string // 在实际应用中，这里应该是更复杂的结构
}

// TXOutput 结构体表示交易的输出
type TXOutput struct {
	Value        int
	ScriptPubKey string // 在实际应用中，这里应该是公钥哈希
}

// SetID 通过序列化交易内容并计算其哈希值来设置交易的ID
func (tx *Transaction) SetID() {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(tx.Vin); err != nil {
		panic(err)
	}
	if err := encoder.Encode(tx.Vout); err != nil {
		panic(err)
	}
	if err := encoder.Encode(tx.Time); err != nil {
		panic(err)
	}

	h := sha256.New()
	h.Write(buffer.Bytes())
	hashed := h.Sum(nil)
	tx.ID = hashed[:]
}

// NewTransaction 创建一个新的交易（这里仍然是一个简化的示例）
func NewTransaction(from, to string, amount int) *Transaction {
	// 注意：这里的 from 和 to 应该是公钥和公钥哈希的字节表示，而不是字符串
	// 为了简化，我们假设它们已经是有效的字节表示，并且跳过签名验证等步骤

	// 创建一个假的交易输入（在实际应用中需要查找UTXO）
	txin := TXInput{[]byte("dummy-txid"), 0, from} // 这里的 Txid 和 Vout 是占位符

	// 创建一个交易输出
	txout := TXOutput{amount, to}

	tx := &Transaction{nil, []TXInput{txin}, []TXOutput{txout}, time.Now().Unix()}
	tx.SetID() // 设置交易的ID

	return tx
}
