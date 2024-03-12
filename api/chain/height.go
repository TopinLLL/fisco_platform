package chain

import (
	"fisco/service/chain"
	"fisco/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetChainHeight
// @Summary      获取区块链高度
// @Tags         管理员操作
// @Router       /blockchain/height [get]
func GetChainHeight(ctx *gin.Context) {
	height := strconv.Itoa(int(chain.GetChainHeight(ctx)))
	response.Success(ctx, "当前区块高度为", height)
}
