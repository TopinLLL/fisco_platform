package generaluser

import (
	"fisco/config"
	"fisco/service/generaluser"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

// ThumbDownData ThumbUpData
// @Summary     数据浏览用户点灭数据
// @Tags         数据浏览用户操作
// @Param		  data_name formData  string  yes "数据名称"
// @Param		  data_content formData  string  yes "数据内容"
// @Router       /data/thumbdown [post]
func ThumbDownData(ctx *gin.Context) {
	dataName := ctx.PostForm("data_name")
	dataContent := ctx.PostForm("data_content")
	session, _ := config.Store.Get(ctx.Request, "sessionID")
	err := generaluser.ThumbDownData(session, dataName, dataContent)
	if err != nil {
		//config.Logger.Error(err.Error())
		response.GeneralFail(ctx, nil, err.Error())
	} else {
		//config.Logger.Info("点灭成功")
		response.Success(ctx, nil, "点灭成功")
	}
}
