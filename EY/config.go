package main

import "time"

// CoinConfig 结构体包含EIYARO硬币的配置信息
type CoinConfig struct {
	Name                 string        // 硬币名字
	Symbol               string        // 硬币符号
	TotalSupply          int64         // 硬币总量
	InitialBlockReward   int64         // 初始区块奖励
	RewardDecreaseRate   float64       // 奖励每年递减率
	TeamRewardPercentage float64       // 团队奖励百分比
	NodeRewardPercentage float64       // 节点奖励百分比
	BlockInterval        time.Duration // 区块间隔时间
}

// NewCoinConfig 创建一个新的EIYARO硬币配置
func NewCoinConfig() *CoinConfig {
	return &CoinConfig{
		Name:                 "EIYARO",
		Symbol:               "EY",
		TotalSupply:          2100000000,      // 21亿
		InitialBlockReward:   1000,            // 初始区块奖励1000 EY
		RewardDecreaseRate:   0.1,             // 每年递减10%
		TeamRewardPercentage: 0.01,            // 团队奖励1%
		NodeRewardPercentage: 0.10,            // 节点奖励10%
		BlockInterval:        3 * time.Minute, // 区块间隔时间为3分钟
	}
}
