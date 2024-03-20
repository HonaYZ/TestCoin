package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
)

// Wallet 结构体表示钱包
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

// NewWallet 生成新的钱包，使用P-256曲线
func NewWallet() (*Wallet, error) {
	// 椭圆曲线实例，P-256
	curve := elliptic.P256()

	// 生成私钥
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	return &Wallet{
		PrivateKey: private,
		PublicKey:  &private.PublicKey,
	}, nil
}

// GetAddress 获取钱包地址（简化版，仅用于演示）
func (w *Wallet) GetAddress() string {
	// 将公钥编码为未压缩的字节形式
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(w.PublicKey)
	if err != nil {
		panic(err) // 在实际应用中，应返回错误而不是panic
	}

	// 计算哈希值
	hash := sha256.Sum256(pubKeyBytes)

	// 将哈希值转换为十六进制字符串
	address := hex.EncodeToString(hash[:])

	return address
}

func main() {
	wallet, err := NewWallet()
	if err != nil {
		fmt.Println("Error creating wallet:", err)
		return
	}
	fmt.Println("Wallet address:", wallet.GetAddress())
	fmt.Println("PrivateKey:", wallet.PrivateKey)
	fmt.Println("PublicKey:", wallet.PublicKey)

}
