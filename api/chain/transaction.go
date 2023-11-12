package chain

import (
	"encoding/hex"
	"fisco/config"
	"fisco/dao/chain/store"
	"fisco/smartcontract/kvtabletest"
	"fmt"
	"math/big"
	"strconv"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

const userPrivateKey = "8632da3eb3256b7f7008fb66a640150d9c1604d96166c9f64fd6477b2fef740f"
const smartContractAddress = "0xC8D6dF8522d2dd192c50514274a46E95E24104a7"

// DeploySmartContract
// @Summary      部署智能合约
// @Tags         链上操作
// @Router       /blockchain/contract/deploy [get]
func DeploySmartContract(ctx *gin.Context) {
	client, err := initSmartContractClient()
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	address, tx, _, err := kvtabletest.DeployKVTableTest(client.GetTransactOpts(), client)
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}

	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())

	// 存储合约地址
	err = store.SmartContract("交易", address.Hex())
	if err != nil {
		config.Logger.Error(err.Error())
	}
}

// SetTransaction
// @Summary      进行交易
// @Tags         链上操作
// @Param		  from formData  string  yes "交易发送方"
// @Param		  to formData  string  yes "交易接收方"
// @Param		  money formData  string  yes "交易金额"
// @Router       /blockchain/contract/set [post]
func SetTransaction(ctx *gin.Context) {
	from := ctx.PostForm("from")
	to := ctx.PostForm("to")
	moneyStr := ctx.PostForm("money")

	// 初始化智能合约客户端
	c, err := initSmartContractClient()
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}

	// 加载智能合约
	contractAddress := common.HexToAddress(smartContractAddress) // deploy contract to get address
	instance, err := kvtabletest.NewKVTableTest(contractAddress, c)
	if err != nil {
		config.Logger.Fatal(err.Error())
		return
	}

	// 进行交易
	session := &kvtabletest.KVTableTestSession{Contract: instance, CallOpts: *c.GetCallOpts(), TransactOpts: *c.GetTransactOpts()}
	money, _ := strconv.Atoi(moneyStr)
	txMoney := big.NewInt(int64(money))
	tx, receipt, err := session.Set(from, txMoney, to)
	if err != nil {
		config.Logger.Fatal(err.Error())
		return
	}

	// 区块信息存入数据库
	hash := common.HexToHash(receipt.BlockHash)
	blockInfo, _ := c.GetBlockByHash(ctx, hash, false)
	if err := store.BlockInfo(blockInfo.ParentHash, blockInfo.Hash, tx.Hash().String(), blockInfo.Number); err != nil {
		config.Logger.Fatal(err.Error())
		return
	}

	// 交易信息存入数据库
	if err := store.Transaction(from, to, int64(money)); err != nil {
		config.Logger.Fatal(err.Error())
		return
	}
	if err := store.UserProperty(from, to, int64(money)); err != nil {
		config.Logger.Fatal(err.Error())
		return
	}
}

// GetDataFromBlockChain
// @Summary      获取链上数据
// @Tags         链上操作
// @Router       /blockchain/contract/get [get]
func GetDataFromBlockChain(ctx *gin.Context) {

}

func initSmartContractClient() (*client.Client, error) {
	privateKey, _ := hex.DecodeString(userPrivateKey)
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

	client, err := client.Dial(c)
	if err != nil {
		return nil, err
	}
	return client, nil
}
