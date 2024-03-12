package test

import (
	"fisco/chameleon"
	"fisco/config"
	"fisco/model"

	"github.com/gin-gonic/gin"
)

// Amend
// @Summary     可修正
// @Tags         实验
// @Router       /test/amend [post]
func Amend(ctx *gin.Context) {
	blockBeforeEdit := &model.TxBlockBeforeEdit{}
	//判断是否修改过
	blockAfterEdit := &model.TxBlockAfterEdit{
		ParentHash:    blockBeforeEdit.ParentHash,
		ParentHashCHA: blockBeforeEdit.ParentHashCHA,
		Height:        blockBeforeEdit.Height,
		TXHash:        blockBeforeEdit.TXHash,
	}
	blockAfterEdit.Hash = chameleon.Seal("test", "test")
	blockAfterEdit.HashCHA = chameleon.Seal("test", "test")[:20]

	if !blockBeforeEdit.HasEdited {
		if err := config.DB.Model(&model.TxBlockBeforeEdit{}).Where("tx_hash=?", "").Update("has_edited", true).Error; err != nil {
		}
		if err := config.DB.Model(&model.TxBlockAfterEdit{}).Create(blockAfterEdit).Error; err != nil {
			config.Logger.Error(err.Error())
		}
	} else {
		if err := config.DB.Model(&model.TxBlockAfterEdit{}).Where("tx_hash=?", "").Updates(blockAfterEdit).Error; err != nil {
		}
	}
	mockAmend(ctx)
	return
}

// AmendAttack
// @Summary     可修正作恶节点
// @Tags         实验
// @Router       /test/amendattack [post]
func AmendAttack(ctx *gin.Context) {
	blockBeforeEdit := &model.TxBlockBeforeEdit{}
	//判断是否修改过
	blockAfterEdit := &model.TxBlockAfterEdit{
		ParentHash:    blockBeforeEdit.ParentHash,
		ParentHashCHA: blockBeforeEdit.ParentHashCHA,
		Height:        blockBeforeEdit.Height,
		TXHash:        blockBeforeEdit.TXHash,
	}
	blockAfterEdit.Hash = chameleon.Seal("test", "test")
	blockAfterEdit.HashCHA = chameleon.Seal("test", "test")[:20]

	if !blockBeforeEdit.HasEdited {
		if err := config.DB.Model(&model.TxBlockBeforeEdit{}).Where("tx_hash=?", "").Update("has_edited", true).Error; err != nil {
		}
		if err := config.DB.Model(&model.TxBlockAfterEdit{}).Create(blockAfterEdit).Error; err != nil {
			config.Logger.Error(err.Error())
		}
	} else {
		if err := config.DB.Model(&model.TxBlockAfterEdit{}).Where("tx_hash=?", "").Updates(blockAfterEdit).Error; err != nil {
		}
	}
	mockAttackAmend(ctx)
	return
}
