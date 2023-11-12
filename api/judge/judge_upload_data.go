package judge

import (
	"fisco/service/judge"
	"fisco/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// JudgeUploadData
// @Summary     审计数据同步至主平台
// @Tags      审计用户操作
// @Param data_id formData string yes "数据ID"
// @Param agree formData boolean yes "该数据是否合法"
// @Router       /data/jugde [post]
func JudgeUploadData(ctx *gin.Context) {
	agree := ctx.PostForm("agree")
	agreeBool, _ := strconv.ParseBool(agree)
	dataID := ctx.PostForm("data_id")
	dataIDNum, _ := strconv.Atoi(dataID)
	err := judge.UploadData(agreeBool, dataIDNum)
	if err != nil {
		//config.Logger.Error(err.Error())
		response.GeneralFail(ctx, nil, err.Error())
	} else {
		//config.Logger.Info("审计数据成功")
		response.Success(ctx, nil, "审计数据成功")
	}
}
