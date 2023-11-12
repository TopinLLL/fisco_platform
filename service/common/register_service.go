package common

import (
	"fisco/dao/common"
	"fisco/model"
	commonUtils "fisco/utils/common"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, username string, password string, roleID int, mail string) (*model.Verify, error) {
	encryptPass, _ := commonUtils.Encrypt(password)
	return common.Register(ctx, username, encryptPass, roleID, mail)
}
