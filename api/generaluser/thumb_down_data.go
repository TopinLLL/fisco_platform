package generaluser

import (
	"fisco/config"
	"fisco/service/generaluser"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

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
