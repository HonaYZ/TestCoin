package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// Transaction 结构体表示交易
type Transaction struct {
	ID          []byte
	Vin         []TXInput
	Vout        []TXOutput
	BlockHash   []byte
	BlockHeight int64
}

// TXInput 表示交易的输入
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

// TXOutput 表示交易的输出
type TXOutput struct {
	Value        int64
	ScriptPubKey string
}

// 假设有一个签名函数，这里简化处理
func sign(privateKey []byte, message []byte) (signature string) {
	// 实际签名实现会涉及私钥和消息的加密处理
	// 这里仅返回一个固定的字符串作为示例
	return "fakeSignature"
}

// 假设有一个验证签名函数，这里简化处理
func verifySignature(publicKey []byte, message []byte, signature string) bool {
	// 实际验证会涉及公钥、消息和签名的解密处理
	// 这里仅检查签名是否与预期匹配
	return signature == "fakeSignature"
}

// NewTransaction 创建交易的方法
func NewTransaction(vin []TXInput, vout []TXOutput) (*Transaction, error) {
	// 实际应用中，这里会有更复杂的逻辑，比如检查输入和输出的金额是否平衡，
	// 检查输入是否真实存在，等等。

	// 为简单起见，我们直接为交易分配一个ID（通常基于交易内容计算哈希值）
	id := sha256.Sum256([]byte("fakeTransactionId"))

	tx := &Transaction{
		ID:          id[:],
		Vin:         vin,
		Vout:        vout,
		BlockHash:   nil, // 交易还未被包含进区块时，BlockHash 为 nil
		BlockHeight: 0,   // 交易还未被包含进区块时，BlockHeight 为 0
	}

	return tx, nil
}

// Sign 签名交易的方法
func (tx *Transaction) Sign(privateKey []byte) error {
	// 签名通常基于交易的一部分内容，例如交易ID和输入/输出列表
	// 这里我们仅对交易ID进行签名作为示例
	message := tx.ID
	for _, vin := range tx.Vin {
		signature := sign(privateKey, message)
		vin.ScriptSig = signature
	}

	return nil
}

// Verify 验证交易的方法
func (tx *Transaction) Verify(publicKey []byte) error {
	// 验证通常检查每个输入上的签名
	for _, vin := range tx.Vin {
		message := tx.ID
		if !verifySignature(publicKey, message, vin.ScriptSig) {
			return errors.New("invalid signature")
		}
	}
	return nil
}

func main() {
	// 示例：创建交易、签名和验证
	vin := []TXInput{
		{Txid: []byte("previousTxId"), Vout: 0, ScriptSig: ""},
	}
	vout := []TXOutput{
		{Value: 100, ScriptPubKey: "recipientPublicKey"},
	}

	tx, err := NewTransaction(vin, vout)
	if err != nil {
		panic(err)
	}

	err = tx.Sign([]byte("senderPrivateKey"))
	if err != nil {
		panic(err)
	}

	err = tx.Verify([]byte("senderPublicKey"))
	if err != nil {
		panic(err)
	}

	// 输出交易ID以验证流程是否成功
	println("Transaction ID:", hex.EncodeToString(tx.ID))
}
