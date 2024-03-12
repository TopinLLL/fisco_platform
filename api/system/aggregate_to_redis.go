package system

import (
	"fisco/service/system"
	"fisco/utils/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AggregateToRedis
// @Summary     聚合密钥
// @Tags         隐私保护模块
// @Param start formData boolean yes "是否开启聚合"
// @Router       /privacy/aggregate [post]
func AggregateToRedis(ctx *gin.Context) {
	start := ctx.PostForm("start")
	startBool, _ := strconv.ParseBool(start)
	//定时任务
	aggregateTicker := time.NewTicker(time.Second * 2)
	if startBool == true {
		for {
			select {
			case <-aggregateTicker.C:
				err := system.AggregateToRedis()
				if err != nil {
					//config.Logger.Error(err.Error())
					response.GeneralFail(ctx, nil, err.Error())
					aggregateTicker.Stop()
					break
				} else {
					//config.Logger.Info("聚合数据成功")
					response.Success(ctx, nil, "聚合数据成功")
				}
			}
		}
	} else {
		aggregateTicker.Stop()
	}
}
