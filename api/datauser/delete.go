package datauser

import (
	"fisco/service/datauser"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

func DeleteData(ctx *gin.Context) {
	dataName := ctx.PostForm("data_name")
	dataContent := ctx.PostForm("data_content")
	err := datauser.DeleteData(ctx, "sessionID", dataName, dataContent)
	if err != nil {
		//config.Logger.Error(err.Error())
		response.GeneralFail(ctx, nil, err.Error())
	} else {
		//config.Logger.Info("数据删除成功")
		response.Success(ctx, nil, "数据删除成功")
	}
}
