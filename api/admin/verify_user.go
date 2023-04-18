package admin

import (
	"fisco/config"
	"fisco/service/admin"
	"fisco/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// VerifyUser
// @Summary      处理用户账号申请
// @Tags         管理员操作
// @Param		  offer_id formData  string  yes "申请id"
// @Param		  ok formData  boolean  yes "是否同意创建"
// @Router       /verify/user [post]
func VerifyUser(ctx *gin.Context) {
	userID := ctx.PostForm("offer_id")
	UserIDNum, _ := strconv.Atoi(userID)
	ok := ctx.PostForm("ok")
	okBool := false
	if ok == "true" {
		okBool = true
	}
	user, err := admin.VerifyUser(UserIDNum, okBool)
	if err != nil {
		response.GeneralFail(ctx, nil, err.Error())
		config.Logger.Error(err.Error())
	} else {
		response.Success(ctx, *user, "已审批用户创建请求")
		config.Logger.Info("已审批用户创建请求")
	}
}
