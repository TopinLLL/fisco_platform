package routes

import (
	"fisco/api/admin"
	"fisco/api/chain"
	"fisco/api/common"
	"fisco/api/contractuser"
	"fisco/api/datauser"
	"fisco/api/generaluser"
	"fisco/api/system"
	"fisco/api/test"
	"fisco/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() (r *gin.Engine) {
	r = gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/login", common.Login)
	r.GET("/blockchain/height", chain.GetChainHeight)
	r.POST("/blockchain/contract/deploy", chain.DeploySmartContract)
	r.POST("/blockchain/contract/singletx", chain.SingleTx)
	r.POST("/blockchain/contract/edittx", chain.EditTX)
	r.POST("/blockchain/contract/tracetx", chain.TraceTX)
	r.POST("/privacy/open", common.PrivacyProtect)

	//使用权限验证中间件
	//r.Use(middleware.NewAuthorizer(casbin.E))
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
	r.POST("/test/amend", test.Amend)
	r.POST("/test/amendattack", test.AmendAttack)
	r.POST("/test/privacy", test.Privacy)
	r.POST("/test/privacydoor", test.PrivacyDoor)

	return
}
