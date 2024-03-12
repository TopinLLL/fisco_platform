package generaluser

import (
	"fisco/service/generaluser"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

func ShowHotData(ctx *gin.Context) {
	results, err := generaluser.ShowHotData()
	if err != nil {
		//config.Logger.Error(err.Error())
		response.GeneralFail(ctx, nil, err.Error())
	} else {
		//config.Logger.Info("查询记录成功")
		response.Success(ctx, results, "查询记录成功")
	}
}
