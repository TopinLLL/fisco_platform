package test

import (
	"context"
	"fisco/chameleon"
	"fisco/config"
	"fisco/model"
	"fisco/utils/response"
	"flag"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "123.57.211.11:9005", "the address to connect to")
)

// Amend
// @Summary     可修正
// @Tags         实验
// @Router       /test/amend [post]
func Amend(ctx *gin.Context) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewTestClient(conn)

	cont, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetTestResult(cont, &AmendableTest{TestNumber: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

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
	mockTime()
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
	response.Success(ctx, nil, res.Message)
	return
}

// AmendAttack
// @Summary     可修正作恶节点
// @Tags         实验
// @Router       /test/amendattack [post]
func AmendAttack(ctx *gin.Context) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewTestClient(conn)

	cont, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetTestResult(cont, &AmendableTest{TestNumber: 2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
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
	mockTime()
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
	response.Success(ctx, nil, res.Message)
	return
}
