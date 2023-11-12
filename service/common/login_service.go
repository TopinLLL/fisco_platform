package common

import (
	"fisco/config"
	"fisco/dao/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, username, password string) (string, error) {
	//数据库中查询是否存在用户
	user, err := common.Login(username, password)
	if err != nil {
		return "", err
	}

	//先尝试获取是否有该session
	session, err := config.Store.Get(ctx.Request, "sessionID")
	if err != nil {
		return "", err
	}

	//防止有session的情况重复登录
	session.Options.MaxAge = -1
	//创建新的session
	session, err = config.Store.New(ctx.Request, "sessionID")
	//给session添加信息
	session.Values["username"] = username
	role, err := common.Role(user.RoleID)
	if err != nil {
		return "", err
	}
	session.Values["role"] = role
	fmt.Println(session.Values["role"])
	session.Options.MaxAge = 86400
	if err != nil {
		return "", err
	}

	//将session存入redis
	err = config.Store.Save(ctx.Request, ctx.Writer, session)
	if err != nil {
		return "", err
	}
	return session.ID, nil
}
