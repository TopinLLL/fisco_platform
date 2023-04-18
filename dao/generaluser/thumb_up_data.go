package generaluser

import (
	"encoding/json"
	"fisco/config"
	"fisco/dto"
	"fisco/global"
	"time"

	"github.com/gorilla/sessions"
)

func ThumbUpData(session *sessions.Session, dataName, dataContent string) error {
	//不允许重复点赞
	username := session.Values["username"].(string)
	nowTime := time.Now().Format("2006-01-02")
	dataDetail, err := config.RDB.HGet(global.BackGround, nowTime+"_data", dataName).Result()
	if err != nil {
		return err
	}
	dataObj := dto.EveryDayData{}
	err = json.Unmarshal([]byte(dataDetail), &dataObj)
	if err != nil {
		return err
	}
	content := dataObj.DataContent[dataContent]
	content.ThumbUpPerson = append(content.ThumbUpPerson, username)
	dataObj.DataContent[dataContent] = dto.DataContentDetail{
		ThumbUp:         content.ThumbUp + 1,
		ThumbDown:       content.ThumbDown,
		ThumbUpPerson:   content.ThumbUpPerson,
		ThumbDownPerson: content.ThumbDownPerson,
	}

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
