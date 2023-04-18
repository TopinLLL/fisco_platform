package datauser

import (
	"encoding/json"
	"fisco/config"
	"fisco/dto"
	"fisco/global"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteData(ctx *gin.Context, sessionID string, dataName string, dataContent string) error {
	//session, _ := config.Store.Get(ctx.Request, sessionID)
	//username := session.Values["username"].(string)
	nowTime := time.Now().Format("2006-01-02")

	data, err := config.RDB.HGet(global.BackGround, nowTime+"_data", dataName).Result()
	if err != nil {
		return err
	}
	dataObj := &dto.EveryDayData{}
	err = json.Unmarshal([]byte(data), dataObj)
	if err != nil {
		return err
	}
	delete(dataObj.DataContent, dataContent)
	dataMarshal, err := json.Marshal(dataObj)
	if err != nil {
		return err
	}
	err = config.RDB.HSet(global.BackGround, nowTime+"_data", dataName, dataMarshal).Err()
	if err != nil {
		return err
	}
	return nil
}
