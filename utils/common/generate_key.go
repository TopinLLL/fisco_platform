package common

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateKey() (string, string, string, error) {
	//生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", "", errors.New("密钥生成失败")
	}
	//将私钥转换成字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//再转换为字符串
	privateKeyString := hexutil.Encode(privateKeyBytes)[2:]
	//将私钥转换为公钥
	publicKey := privateKey.Public()
	//判断是否为ecdsa类型公钥
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//将公钥转换为字节
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//将公钥转换为字符串
	publicKeyString := hexutil.Encode(publicKeyBytes)[4:]
	//将公钥转换为地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return publicKeyString, privateKeyString, address, nil
}
