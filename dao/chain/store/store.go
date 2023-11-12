package store

import (
	"fisco/config"
	"fisco/model"
)

// SmartContract 存储智能合约
func SmartContract(name string, address string) error {
	contract := &model.Contract{
		ContractName:    name,
		ContractAddress: address,
	}

	return config.DB.Model(&model.Contract{}).Create(contract).Error
}

// BlockInfo 存储区块信息
func BlockInfo(parentHash, hash, txHash, height string) error {
	block := &model.Block{
		ParentHash: parentHash,
		Hash:       hash,
		Height:     height,
		TXHash:     txHash,
	}
	return config.DB.Model(&model.Block{}).Create(block).Error
}

// Transaction 存储交易信息
func Transaction(from, to string, money int64) error {
	tx := &model.Transaction{
		From:  from,
		Money: money,
		To:    to,
	}

	return config.DB.Model(&model.Transaction{}).Create(tx).Error
}

// UserProperty 存储用户财产信息
func UserProperty(from, to string, money int64) error {
	user := &model.UserProperty{}

	config.DB.Model(&model.UserProperty{}).Where("username=?", from).Find(user)
	fromPropertyLeft := user.Property - money
	if err := config.DB.Model(&model.UserProperty{}).Where("username=?", from).Update("property", fromPropertyLeft).Error; err != nil {
		return err
	}

	config.DB.Model(&model.UserProperty{}).Where("username=?", to).Find(user)
	toPropertyLeft := user.Property + money
	if err := config.DB.Model(&model.UserProperty{}).Where("username=?", to).Update("property", toPropertyLeft).Error; err != nil {
		return err
	}

	return nil
}
