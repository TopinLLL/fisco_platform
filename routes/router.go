package routes

import (
	"fisco/api/admin"
	"fisco/api/common"
	"fisco/api/contractuser"
	"fisco/api/datauser"
	"fisco/api/generaluser"
	"fisco/api/judge"
	"fisco/api/system"
	"fisco/casbin"
	"fisco/docs"
	"fisco/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() (r *gin.Engine) {
	r = gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/login", common.Login)
	//使用权限验证中间件
	r.Use(middleware.NewAuthorizer(casbin.E))
	r.POST("/add/role", admin.AddRole)
	r.POST("/register", common.Register)
	r.POST("/verify/user", admin.VerifyUser)
	r.POST("/verify/email", contractuser.Verify)
	r.POST("/data/provide", datauser.ProvideData)
	r.POST("/data/delete", datauser.DeleteData)
	r.GET("/show/hot", generaluser.ShowHotData)
	r.POST("/data/thumbup", generaluser.ThumbUpData)
	r.POST("/data/thumbdown", generaluser.ThumbDownData)
	r.POST("/system/aggregate", system.AggregateToRedis)
	r.POST("/data/jugde", judge.JudgeUploadData)
	return
}
