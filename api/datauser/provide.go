package datauser

import (
	"fisco/service/datauser"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

func ProvideData(ctx *gin.Context) {
	dataName := ctx.PostForm("data_name")
	dataContent := ctx.PostForm("data_content")
	err := datauser.ProvideData(ctx, "sessionID", dataName, dataContent)
	if err != nil {
		//config.Logger.Error(err.Error())
		response.GeneralFail(ctx, nil, err.Error())
	} else {
		//config.Logger.Info("数据提供成功")
		response.Success(ctx, nil, "数据提供成功")
	}
}
