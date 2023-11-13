package store

import (
	"fisco/chameleon"
	"fisco/config"
	"fisco/model"
	"fmt"
	"strconv"

	"github.com/FISCO-BCOS/go-sdk/core/types"
)

// SmartContract 存储智能合约
func SmartContract(name string, address string) error {
	contract := &model.Contract{
		ContractName:    name,
		ContractAddress: address,
	}

	return config.DB.Model(&model.Contract{}).Create(contract).Error
}

// BlockInfoBeforeEdit 存储编辑前区块信息
func BlockInfoBeforeEdit(blockInfo *types.Block, privateKey, txHash string, data string) error {
	parentHash := chameleon.Chameleon(fmt.Sprintf("%s%s", blockInfo.ParentHash, privateKey), data)
	hash := chameleon.Chameleon(fmt.Sprintf("%s%s", blockInfo.Hash, privateKey), data)

	block := &model.TxBlockBeforeEdit{
		ParentHash:    parentHash,
		ParentHashCHA: parentHash[:20],
		Hash:          hash,
		HashCHA:       hash[:20],
		Height:        blockInfo.Number,
		TXHash:        txHash,
	}
	return config.DB.Model(&model.TxBlockBeforeEdit{}).Create(block).Error
}

// Transaction 存储交易信息
func Transaction(from, to string, money int64, txBlockHeight string, txHash string) error {
	tx := &model.TxDealDetail{
		From:          from,
		Money:         money,
		To:            to,
		TxBlockHeight: txBlockHeight,
		TxHash:        txHash,
	}

	return config.DB.Model(&model.TxDealDetail{}).Create(tx).Error
}

// UserProperty 存储用户财产信息
func UserProperty(from, to string, money int64) error {
	userFrom := &model.TxUserProperty{}

	config.DB.Model(&model.TxUserProperty{}).Where("username=?", from).Find(userFrom)
	fromPropertyLeft := userFrom.Property - money
	if err := config.DB.Model(&model.TxUserProperty{}).Where("username=?", from).Update("property", fromPropertyLeft).Error; err != nil {
		return err
	}

	userTo := &model.TxUserProperty{}
	config.DB.Model(&model.TxUserProperty{}).Where("username=?", to).Find(userTo)
	toPropertyLeft := userTo.Property + money
	if err := config.DB.Model(&model.TxUserProperty{}).Where("username=?", to).Update("property", toPropertyLeft).Error; err != nil {
		return err
	}

	return nil
}

// EditTX 编辑
func EditTX(privateKey string, txHash string, data string) error {
	blockBeforeEdit := &model.TxBlockBeforeEdit{}
	//判断是否修改过
	if err := config.DB.Model(&model.TxBlockBeforeEdit{}).Where("tx_hash=?", txHash).Find(blockBeforeEdit).Error; err != nil {
		return err
	}

	blockAfterEdit := &model.TxBlockAfterEdit{
		ParentHash:    blockBeforeEdit.ParentHash,
		ParentHashCHA: blockBeforeEdit.ParentHashCHA,
		Height:        blockBeforeEdit.Height,
		TXHash:        blockBeforeEdit.TXHash,
	}
	blockAfterEdit.Hash = chameleon.Chameleon(fmt.Sprintf("%s%s", blockBeforeEdit.Hash, privateKey), data)
	blockAfterEdit.HashCHA = chameleon.Chameleon(privateKey, data)[:20]

	if !blockBeforeEdit.HasEdited {
		if err := config.DB.Model(&model.TxBlockBeforeEdit{}).Where("tx_hash=?", txHash).Update("has_edited", true).Error; err != nil {
			return err
		}
		if err := config.DB.Model(&model.TxBlockAfterEdit{}).Create(blockAfterEdit).Error; err != nil {
			config.Logger.Error(err.Error())
			return err
		}
	} else {
		if err := config.DB.Model(&model.TxBlockAfterEdit{}).Where("tx_hash=?", txHash).Updates(blockAfterEdit).Error; err != nil {
			return err
		}
	}
	// 编辑用户资产
	if err := editUserProperty(txHash, data); err != nil {
		return err
	}
	// 编辑交易
	if err := editTXDetail(txHash, data); err != nil {
		return err
	}
	return nil
}

func editTXDetail(txHash, data string) error {
	dealDetail := &model.TxDealDetail{}
	if err := config.DB.Model(&model.TxDealDetail{}).Where("tx_hash=?", txHash).Find(dealDetail).Error; err != nil {
		return err
	}

	editMoney, _ := strconv.Atoi(data)
	if err := config.DB.Model(&model.TxDealDetail{}).Where("tx_hash=?", txHash).Update("money", int64(editMoney)).Error; err != nil {
		config.Logger.Error(err.Error())
		return err
	}

	return nil
}

func editUserProperty(txHash, data string) error {
	dealDetail := &model.TxDealDetail{}
	if err := config.DB.Model(&model.TxDealDetail{}).Where("tx_hash=?", txHash).Find(dealDetail).Error; err != nil {
		return err
	}

	from := dealDetail.From
	to := dealDetail.To
	moneyBefore := dealDetail.Money
	moneyNow, _ := strconv.Atoi(data)
	moneyDiff := int64(moneyNow) - moneyBefore

	FromUserProperty := &model.TxUserProperty{}
	if err := config.DB.Model(&model.TxUserProperty{}).Where("username=?", from).Find(FromUserProperty).Error; err != nil {
		return err
	}
	fromUserPropertyBefore := FromUserProperty.Property
	fromUserPropertyNow := fromUserPropertyBefore + moneyDiff
	if err := config.DB.Model(&model.TxUserProperty{}).Where("username=?", from).Update("property", fromUserPropertyNow).Error; err != nil {
		return err
	}

	ToUserProperty := &model.TxUserProperty{}
	if err := config.DB.Model(&model.TxUserProperty{}).Where("username=?", to).Find(ToUserProperty).Error; err != nil {
		return err
	}
	toUserPropertyBefore := ToUserProperty.Property
	toUserPropertyNow := toUserPropertyBefore - moneyDiff
	if err := config.DB.Model(&model.TxUserProperty{}).Where("username=?", to).Update("property", toUserPropertyNow).Error; err != nil {
		return err
	}
	return nil
}
