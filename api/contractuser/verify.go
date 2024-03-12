package contractuser

import (
	"fisco/config"
	"fisco/service/contractuser"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

// Verify
// @Summary      确认智能合约
// @Tags         管理员操作
// @Param		  verify_email formData  string  yes "验证邮件"
// @Router       /verify/smartcontract [post]
func Verify(ctx *gin.Context) {
	session, _ := config.Store.Get(ctx.Request, "sessionID")
	username := session.Values["username"]
	verifyEmail := ctx.PostForm("verify_email")
	err := contractuser.Verify(username.(string), verifyEmail)
	if err != nil {
		response.GeneralFail(ctx, nil, err.Error())
		//config.Logger.Error(err.Error())
	} else {
		response.Success(ctx, nil, "确认成功")
		//config.Logger.Info("确认成功")
	}
}
