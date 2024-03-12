package chain

import (
	"fisco/config"
	"fisco/dao/chain/store"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

// TraceTX
// @Summary      追溯修正操作
// @Tags         电子合同通信与修正模块
// @Param		 privateKey formData  string  yes "监管者私钥"
// @Param		 txHash formData  string  yes "交易地址"
// @Router       /amend/contract/tracetx [post]
func TraceTX(ctx *gin.Context) {
	txHash := ctx.PostForm("txHash")
	detail, err := store.TraceTX(txHash)
	if err != nil {
		config.Logger.Fatal(err.Error())
		return
	}

	response.Success(ctx, detail, "")
}
