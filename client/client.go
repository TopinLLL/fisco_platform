package client

import (
	"context"
	"encoding/hex"
	"fisco/config"
	"fmt"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

func Client() {
	//将私钥解码成.pem文件
	privateKey, err := hex.DecodeString("8632da3eb3256b7f7008fb66a640150d9c1604d96166c9f64fd6477b2fef740f")

	if err != nil {
		config.Logger.Error(err.Error())
	}

	con := conf.Config{
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
	client, err := client.Dial(&con)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	number, err := client.GetBlockNumber(context.Background())
	//world, transaction, h, err := helloworld.DeployHelloWorld(client.GetTransactOpts(), client)
	//if err != nil {
	//	config.Logger.Error(err.Error())
	//}
	fmt.Println(number)
}
