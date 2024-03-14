package common

import (
	"encoding/json"
	"fisco/config"
	"fisco/dto"
	"fisco/global"
	"fisco/utils/response"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type TeamMember struct {
	Name        string
	Position    string
	ContactInfo string
}

var teamMembers []TeamMember

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
	signature := generateGroupSignature(teamMembers)
	fmt.Println("群签名：")
	mockTime()
	for position, members := range signature {
		fmt.Printf("%s：\n", position)
		for _, member := range members {
			fmt.Printf("过程：%s，%s\n", member.Name, member.ContactInfo)
		}
		dataContent := ""
		dataName := ""
		session := &sessions.Session{}
		username := session.Values["username"].(string)
		nowTime := time.Now().Format("2006-01-02")
		dataDetail, _ := config.RDB.HGet(global.BackGround, nowTime+"_data", dataName).Result()
		dataObj := dto.EveryDayData{}
		_ = json.Unmarshal([]byte(dataDetail), &dataObj)
		content := dataObj.DataContent[dataContent]
		content.ThumbDownPerson = append(content.ThumbDownPerson, username)
		dataObj.DataContent[dataContent] = dto.DataContentDetail{
			ThumbUp:         content.ThumbUp,
			ThumbDown:       content.ThumbDown + 1,
			ThumbUpPerson:   content.ThumbUpPerson,
			ThumbDownPerson: content.ThumbDownPerson,
		}
	}
}

// GroupJoin
// @Summary      群加入
// @Tags         隐私保护模块
// @Param		 signature formData  string  yes "用户签名"
// @Router       /privacy/groupjoin [post]
func GroupJoin(ctx *gin.Context) {
	signature := generateGroupSignature(teamMembers)
	fmt.Println("群签名：")
	mockTime()
	for position, members := range signature {
		fmt.Printf("%s：\n", position)
		for _, member := range members {
			fmt.Printf("过程：%s，%s\n", member.Name, member.ContactInfo)
		}
		dataContent := ""
		dataName := ""
		session := &sessions.Session{}
		username := session.Values["username"].(string)
		nowTime := time.Now().Format("2006-01-02")
		dataDetail, _ := config.RDB.HGet(global.BackGround, nowTime+"_data", dataName).Result()
		dataObj := dto.EveryDayData{}
		_ = json.Unmarshal([]byte(dataDetail), &dataObj)
		content := dataObj.DataContent[dataContent]
		content.ThumbDownPerson = append(content.ThumbDownPerson, username)
		dataObj.DataContent[dataContent] = dto.DataContentDetail{
			ThumbUp:         content.ThumbUp,
			ThumbDown:       content.ThumbDown + 1,
			ThumbUpPerson:   content.ThumbUpPerson,
			ThumbDownPerson: content.ThumbDownPerson,
		}
	}
}

// ContractSign
// @Summary      电子合同签名
// @Tags         隐私保护模块
// @Param		 signature formData  string  yes "用户签名"
// @Router       /privacy/sign [post]
func ContractSign(ctx *gin.Context) {
	signature := generateGroupSignature(teamMembers)
	fmt.Println("群签名：")
	mockTime()
	for position, members := range signature {
		fmt.Printf("%s：\n", position)
		for _, member := range members {
			fmt.Printf("过程：%s，%s\n", member.Name, member.ContactInfo)
		}
		dataContent := ""
		dataName := ""
		session := &sessions.Session{}
		username := session.Values["username"].(string)
		nowTime := time.Now().Format("2006-01-02")
		dataDetail, _ := config.RDB.HGet(global.BackGround, nowTime+"_data", dataName).Result()
		dataObj := dto.EveryDayData{}
		_ = json.Unmarshal([]byte(dataDetail), &dataObj)
		content := dataObj.DataContent[dataContent]
		content.ThumbDownPerson = append(content.ThumbDownPerson, username)
		dataObj.DataContent[dataContent] = dto.DataContentDetail{
			ThumbUp:         content.ThumbUp,
			ThumbDown:       content.ThumbDown + 1,
			ThumbUpPerson:   content.ThumbUpPerson,
			ThumbDownPerson: content.ThumbDownPerson,
		}
	}
}

// SignVerify
// @Summary      群签名验证
// @Tags         隐私保护模块
// @Param		 signature formData  string  yes "用户签名"
// @Router       /privacy/verify [post]
func SignVerify(ctx *gin.Context) {
	signature := generateGroupSignature(teamMembers)
	fmt.Println("群签名：")
	mockTime()
	for position, members := range signature {
		fmt.Printf("%s：\n", position)
		for _, member := range members {
			fmt.Printf("过程：%s，%s\n", member.Name, member.ContactInfo)
		}
		dataContent := ""
		dataName := ""
		session := &sessions.Session{}
		username := session.Values["username"].(string)
		nowTime := time.Now().Format("2006-01-02")
		dataDetail, _ := config.RDB.HGet(global.BackGround, nowTime+"_data", dataName).Result()
		dataObj := dto.EveryDayData{}
		_ = json.Unmarshal([]byte(dataDetail), &dataObj)
		content := dataObj.DataContent[dataContent]
		content.ThumbDownPerson = append(content.ThumbDownPerson, username)
		dataObj.DataContent[dataContent] = dto.DataContentDetail{
			ThumbUp:         content.ThumbUp,
			ThumbDown:       content.ThumbDown + 1,
			ThumbUpPerson:   content.ThumbUpPerson,
			ThumbDownPerson: content.ThumbDownPerson,
		}
	}
}

func generateGroupSignature(members []TeamMember) map[string][]TeamMember {
	signature := make(map[string][]TeamMember)
	for _, member := range members {
		signature[member.Position] = append(signature[member.Position], member)
	}
	return signature
}
