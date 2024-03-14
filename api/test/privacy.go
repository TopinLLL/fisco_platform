package test

import (
	"context"
	"encoding/json"
	"fisco/config"
	"fisco/dto"
	"fisco/global"
	"fisco/utils/response"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	mockTime()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewTestClient(conn)

	cont, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetTestResult(cont, &AmendableTest{TestNumber: 4})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	response.Success(ctx, nil, res.Message)
}

// PrivacyDoor
// @Summary     不同实验群签名
// @Tags         实验
// @Router       /test/privacydoor [post]
func PrivacyDoor(ctx *gin.Context) {
	var teamMembers []TeamMember

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
	mockTime()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewTestClient(conn)

	cont, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetTestResult(cont, &AmendableTest{TestNumber: 3})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	response.Success(ctx, nil, res.Message)
}
