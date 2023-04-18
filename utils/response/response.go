package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS      = http.StatusOK
	DATABASEFAIL = 405
	GENERALFAIL  = http.StatusInternalServerError
)

func Response(ctx *gin.Context, data interface{}, code int, msg string) {
	ctx.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, data, SUCCESS, msg)
}

func GeneralFail(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, data, GENERALFAIL, msg)
}

func DataBaseFail(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, data, DATABASEFAIL, msg)
}
