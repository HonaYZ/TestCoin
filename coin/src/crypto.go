package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// 计算数据的SHA256哈希值
func CalculateHash(data []byte) string {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
