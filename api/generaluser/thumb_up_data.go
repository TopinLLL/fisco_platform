package generaluser

import (
	"fisco/config"
	"fisco/service/generaluser"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

// ThumbUpData
// @Summary     数据浏览用户点赞热点数据
// @Tags         数据浏览用户操作
// @Param		  data_name formData  string  yes "数据名称"
// @Param		  data_content formData  string  yes "数据内容"
// @Router       /data/thumbup [post]
func ThumbUpData(ctx *gin.Context) {
	dataName := ctx.PostForm("data_name")
	dataContent := ctx.PostForm("data_content")
	session, _ := config.Store.Get(ctx.Request, "sessionID")
	err := generaluser.ThumbUpData(session, dataName, dataContent)
	if err != nil {
		//config.Logger.Error(err.Error())
		response.GeneralFail(ctx, nil, err.Error())
	} else {
		//config.Logger.Info("点赞成功")
		response.Success(ctx, nil, "点赞成功")
	}
}
