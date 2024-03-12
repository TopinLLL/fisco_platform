package test

import (
	"encoding/json"
	"fisco/config"
	"fisco/dto"
	"fisco/global"
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

// Privacy
// @Summary     群签名
// @Tags         实验
// @Router       /test/privacy [post]
func Privacy(ctx *gin.Context) {
	var teamMembers []TeamMember

	signature := generateGroupSignature(teamMembers)
	fmt.Println("群签名：")
	time.Sleep(time.Second * 7)
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
	mockPrivacy(ctx)
}

// PrivacyDoor
// @Summary     不同实验群签名
// @Tags         实验
// @Router       /test/privacydoor [post]
func PrivacyDoor(ctx *gin.Context) {
	var teamMembers []TeamMember

	signature := generateGroupSignature(teamMembers)
	fmt.Println("群签名：")
	time.Sleep(time.Second * 7)
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
	mockPrivacyDoor(ctx)
}
