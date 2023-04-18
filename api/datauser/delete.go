package datauser

import (
	"fisco/config"
	"fisco/service/datauser"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

// DeleteData
// @Summary     数据提供用户删除数据
// @Tags         数据提供用户操作
// @Param		  data_name formData  string  yes "数据名称"
// @Param		  data_content formData  string  yes "数据内容"
// @Router       /data/delete [post]
func DeleteData(ctx *gin.Context) {
	dataName := ctx.PostForm("data_name")
	dataContent := ctx.PostForm("data_content")
	err := datauser.DeleteData(ctx, "sessionID", dataName, dataContent)
	if err != nil {
		config.Logger.Error(err.Error())
		response.GeneralFail(ctx, nil, err.Error())
	} else {
		config.Logger.Info("数据删除成功")
		response.Success(ctx, nil, "数据删除成功")
	}
}
