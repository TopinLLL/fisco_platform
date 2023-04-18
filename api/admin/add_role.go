package admin

import (
	"fisco/config"
	"fisco/service/admin"
	"fisco/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddRole
// @Summary      添加用户
// @Tags         管理员操作
// @Param		  role_id formData  string  yes "角色id"
// @Param		  role_name formData  string  yes "角色名称"
// @Router       /add/role [post]
func AddRole(ctx *gin.Context) {
	roleID := ctx.PostForm("role_id")
	roleIDNum, _ := strconv.Atoi(roleID)
	roleName := ctx.PostForm("role_name")
	role, err := admin.AddRole(roleIDNum, roleName)
	if err != nil {
		response.GeneralFail(ctx, nil, err.Error())
		config.Logger.Error(err.Error())
	} else {
		response.Success(ctx, *role, "新增用户种类成功")
		config.Logger.Info("新增用户种类成功")
	}
}
