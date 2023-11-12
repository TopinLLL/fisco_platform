package chain

import (
	"context"
	"encoding/hex"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/gin-gonic/gin"
)

// GetChainHeight 获取区块当前高度
func GetChainHeight(ctx *gin.Context) int64 {
	//将私钥解码成.pem文件
	privateKey, _ := hex.DecodeString("8632da3eb3256b7f7008fb66a640150d9c1604d96166c9f64fd6477b2fef740f")

	c := &conf.Config{
		//是否国密
		IsSMCrypto: false,
		GroupID:    1,
		ChainID:    1,
		//用户私钥
		PrivateKey: privateKey,
		CAFile:     "./keyboard/ca.crt",
		Key:        "./keyboard/sdk.key",
		Cert:       "./keyboard/sdk.crt",
		NodeURL:    "123.57.211.11:20200",
	}
	client, _ := client.Dial(c)

	number, _ := client.GetBlockNumber(context.Background())
	return number
}
