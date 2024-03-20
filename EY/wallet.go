package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
)

// Wallet 结构体表示一个钱包，包含公钥和私钥
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

// NewWallet 创建一个新的钱包，并生成公钥和私钥
func NewWallet() *Wallet {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return nil
	}

	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	wallet := &Wallet{PrivateKey: privateKey, PublicKey: publicKey}
	return wallet
}

// GetAddress 获取钱包的地址（通常是公钥的哈希）
func (w *Wallet) GetAddress() string {
	// 注意：这里为了简化，我们直接将公钥转换为十六进制字符串作为地址。
	// 在实际应用中，你应该使用某种哈希函数（如RIPEMD-160或Keccak-256）来处理公钥，
	// 并可能加上一些网络标识符或校验和来生成最终的地址。
	return hex.EncodeToString(w.PublicKey)
}

func main() {
	// 创建一个新的钱包
	wallet := NewWallet()

	// 打印钱包的地址
	fmt.Printf("Wallet Address: %s\n", wallet.GetAddress())

	// 注意：在实际应用中，你不应该直接打印或存储私钥。
	// 这里仅为了演示目的而打印私钥。
	privateKeyBytes, err := x509.MarshalECPrivateKey(wallet.PrivateKey)
	if err != nil {
		fmt.Println("Error marshaling private key:", err)
		return
	}
	fmt.Printf("Private Key: %x\n", privateKeyBytes)

	// 同样地，为了安全起见，通常也不会直接打印公钥的原始字节。
	// 这里只是为了演示如何访问它。
	fmt.Printf("Public Key: %x\n", wallet.PublicKey)
}
