package main

import "fmt"

// Config 定义区块链的配置参数
type Config struct {
	CoinName       string  // 硬币名字
	CoinSymbol     string  // 硬币符号
	TotalCoins     int64   // 硬币总量
	InitialReward  int64   // 初始区块奖励
	RewardDecay    float64 // 区块奖励递减率
	TeamAllocation float64 // 团队地址分配比例
	NodeAllocation float64 // 节点地址分配比例
	BlockInterval  int64   // 区块间隔时间（秒）
}

// 全局变量使用不同的名称以避免与类型名称冲突
var defaultConfig = Config{
	CoinName:       "EIYARO",
	CoinSymbol:     "EY",
	TotalCoins:     2100000000,
	InitialReward:  1000,
	RewardDecay:    0.1,
	TeamAllocation: 0.01,
	NodeAllocation: 0.1,
	BlockInterval:  3 * 60, // 3分钟
}

// 在其他地方，如果需要引用这个默认配置，请使用 defaultConfig 而不是 Config
func main() {

	fmt.Println(defaultConfig)
}
