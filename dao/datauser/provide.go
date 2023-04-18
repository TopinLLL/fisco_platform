package datauser

import (
	"encoding/json"
	"fisco/config"
	"fisco/dto"
	"fisco/global"
	"time"

	"github.com/gin-gonic/gin"
)

func ProvideData(ctx *gin.Context, sessionID string, dataName string, dataContent string) error {
	session, _ := config.Store.Get(ctx.Request, sessionID)
	username := session.Values["username"].(string)
	nowTime := time.Now().Format("2006-01-02")
	result, err := config.RDB.Keys(global.BackGround, nowTime+"_data").Result()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		detail := dto.DataContentDetail{
			ThumbUp:   0,
			ThumbDown: 0,
		}

		today := dto.EveryDayData{
			DataContent:   make(map[string]dto.DataContentDetail),
			ProvidePerson: username,
		}
		today.DataContent[dataContent] = detail
		dataMarshal, err := json.Marshal(today)
		if err != nil {
			return err
		}
		err = config.RDB.HSet(global.BackGround, nowTime+"_data", dataName, dataMarshal).Err()
		if err != nil {
			return err
		}
		config.RDB.Expire(global.BackGround, nowTime+"_data", time.Hour*24)
	} else {
		hKeys, err := config.RDB.HKeys(global.BackGround, nowTime+"_data").Result()
		if err != nil {
			return err
		}
		var i int
		for _, key := range hKeys {
			if key == dataName {
				break
			}
			i++
		}
		detail := dto.DataContentDetail{
			ThumbUp:   0,
			ThumbDown: 0,
		}
		if i == len(hKeys) {
			today := dto.EveryDayData{
				DataContent:   make(map[string]dto.DataContentDetail),
				ProvidePerson: username,
			}
			today.DataContent[dataContent] = detail
			dataMarshal, err := json.Marshal(today)
			if err != nil {
				return err
			}
			err = config.RDB.HSet(global.BackGround, nowTime+"_data", dataName, dataMarshal).Err()
			if err != nil {
				return err
			}
		} else {
			todayData, err := config.RDB.HGet(global.BackGround, nowTime+"_data", dataName).Result()
			todayObj := &dto.EveryDayData{}
			err = json.Unmarshal([]byte(todayData), todayObj)
			if err != nil {
				return err
			}
			todayObj.DataContent[dataContent] = dto.DataContentDetail{}
			NewMarshal, err := json.Marshal(todayObj)
			if err != nil {
				return err
			}
			err = config.RDB.HSet(global.BackGround, nowTime+"_data", dataName, NewMarshal).Err()
			if err != nil {
				return err
			}
		}

	}

	return nil
}
