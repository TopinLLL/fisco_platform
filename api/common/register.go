package common

import (
	"fisco/service/common"
	"fisco/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Register
// @Summary      用户注册
// @Tags         常规用户操作
// @Param		  username formData  string  yes "账号"
// @Param		  password formData  string  yes "密码"
// @Param		  role_id formData  string  yes "角色id"
// @Param		  mail formData  string  yes "邮箱"
// @Router       /register [post]
func Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	roleID := ctx.PostForm("role_id")
	roleIDNum, _ := strconv.Atoi(roleID)
	mail := ctx.PostForm("mail")
	verify, err := common.Register(ctx, username, password, roleIDNum, mail)
	if err != nil {
		response.GeneralFail(ctx, nil, err.Error())
		//config.Logger.Error(err.Error())
	} else {
		response.Success(ctx, *verify, "已提交用户申请")
		//config.Logger.Info("已提交用户申请")
	}
}
