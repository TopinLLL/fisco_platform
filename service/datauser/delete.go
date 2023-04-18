package datauser

import (
	"fisco/dao/datauser"

	"github.com/gin-gonic/gin"
)

func DeleteData(ctx *gin.Context, sessionID string, dataName, dataContent string) error {
	return datauser.DeleteData(ctx, sessionID, dataName, dataContent)
}
