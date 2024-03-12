package common

import (
	"fisco/utils/response"

	"github.com/gin-gonic/gin"
)

// PrivacyProtect
// @Summary      群签名打开
// @Tags         隐私保护模块
// @Param		 signature formData  string  yes "用户签名"
// @Router       /privacy/open [post]
func PrivacyProtect(ctx *gin.Context) {
	ctx.PostForm("signature")
	response.Success(ctx, gin.H{"用户签名信息": "i am Alice"}, "验证成功")
}

// GroupCreate
// @Summary      群创建
// @Tags         隐私保护模块
// @Param		 signature formData  string  yes "用户签名"
// @Router       /privacy/groupcreate [post]
func GroupCreate(ctx *gin.Context) {

}

// GroupJoin
// @Summary      群加入
// @Tags         隐私保护模块
// @Param		 signature formData  string  yes "用户签名"
// @Router       /privacy/groupjoin [post]
func GroupJoin(ctx *gin.Context) {

}

// ContractSign
// @Summary      电子合同签名
// @Tags         隐私保护模块
// @Param		 signature formData  string  yes "用户签名"
// @Router       /privacy/sign [post]
func ContractSign(ctx *gin.Context) {

}

// SignVerify
// @Summary      群签名验证
// @Tags         隐私保护模块
// @Param		 signature formData  string  yes "用户签名"
// @Router       /privacy/verify [post]
func SignVerify(ctx *gin.Context) {

}
