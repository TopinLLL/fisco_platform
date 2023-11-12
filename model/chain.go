package model

import "gorm.io/gorm"

// Contract 智能合约
type Contract struct {
	*gorm.Model
	ContractName    string `json:"contract_name,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"` // 合约地址
}

// Block 区块信息
type Block struct {
	*gorm.Model
	ParentHash string
	Hash       string
	Height     string
	TXHash     string
}

type Transaction struct {
	*gorm.Model
	From  string
	Money int64
	To    string
}

type UserProperty struct {
	*gorm.Model
	Username string
	Property int64
}
