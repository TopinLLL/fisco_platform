package model

import "gorm.io/gorm"

// Contract 智能合约
type Contract struct {
	*gorm.Model
	ContractName    string `json:"contract_name,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"` // 合约地址
}

// TxBlockBeforeEdit 编辑前区块信息
type TxBlockBeforeEdit struct {
	*gorm.Model
	ParentHash    string
	ParentHashCHA string
	Hash          string
	HashCHA       string
	Height        string
	TXHash        string
	HasEdited     bool
}

// TxBlockAfterEdit 编辑后区块信息
type TxBlockAfterEdit struct {
	*gorm.Model
	ParentHash    string
	ParentHashCHA string
	Hash          string
	HashCHA       string
	Height        string
	TXHash        string
}

type TxDealDetail struct {
	*gorm.Model
	From          string
	Money         int64
	To            string
	TxBlockHeight string
	TxHash        string
}

type TxUserProperty struct {
	*gorm.Model
	Username string
	Property int64
}
