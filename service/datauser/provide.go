package datauser

import (
	"fisco/dao/datauser"

	"github.com/gin-gonic/gin"
)

func ProvideData(ctx *gin.Context, sessionID string, dataName string, dataContent string) error {
	return datauser.ProvideData(ctx, sessionID, dataName, dataContent)
}
