package contractuser

import (
	"fisco/config"
	"fisco/service/contractuser"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

// Verify
// @Summary      智能合约用户确认邮件
// @Tags         智能合约用户操作
// @Param		  verify_email formData  string  yes "验证邮件"
// @Router       /verify/email [post]
func Verify(ctx *gin.Context) {
	session, _ := config.Store.Get(ctx.Request, "sessionID")
	username := session.Values["username"]
	verifyEmail := ctx.PostForm("verify_email")
	err := contractuser.Verify(username.(string), verifyEmail)
	if err != nil {
		response.GeneralFail(ctx, nil, err.Error())
		config.Logger.Error(err.Error())
	} else {
		response.Success(ctx, nil, "确认成功")
		config.Logger.Info("确认成功")
	}
}
