package common

import (
	"fisco/service/common"
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

// Login
// @Summary      用户登录
// @Tags         登录
// @Param		  username formData  string  yes "账号"
// @Param		  password formData  string  yes "密码"
// @Router       /login [post]
func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	sessionID, err := common.Login(ctx, username, password)
	if err != nil {
		response.DataBaseFail(ctx, nil, err.Error())
		//config.Logger.Error(err.Error())
	} else {
		response.Success(ctx, gin.H{"sessionID": sessionID}, "登录成功")
		//config.Logger.Info("登录成功")
	}
}
