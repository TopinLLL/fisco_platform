package system

import (
	"encoding/json"
	"errors"
	"fisco/config"
	"fisco/dto"
	"fisco/global"
	"fisco/model"
	"time"
)

func AggregateToRedis() error {
	timeNow := time.Now().Format("2006-01-02")
	//每次访问Key都要进行错误验证
	todayData, err := config.RDB.Keys(global.BackGround, timeNow+"_data").Result()
	if err != nil {
		return err
	}
	if len(todayData) == 0 {
		return errors.New("暂无可聚合数据，请暂时关闭此功能")
	}

	dataDetailKeys, err := config.RDB.HKeys(global.BackGround, timeNow+"_data").Result()
	if err != nil {
		return err
	}
	for _, field := range dataDetailKeys {
		dataDetail, err := config.RDB.HGet(global.BackGround, timeNow+"_data", field).Result()
		dataDetailObj := dto.EveryDayData{}
		err = json.Unmarshal([]byte(dataDetail), &dataDetailObj)
		if err != nil {
			return err
		}
		//循环每个field中的content
		for contentName := range dataDetailObj.DataContent {
			//同步到Mysql
			thumbNumber := dataDetailObj.DataContent[contentName].ThumbUp - dataDetailObj.DataContent[contentName].ThumbDown
			if thumbNumber >= 5 {
				//在redis中删除该条
				record := dataDetailObj.DataContent[contentName]
				delete(dataDetailObj.DataContent, contentName)
				if len(dataDetailObj.DataContent) == 0 {
					config.RDB.HDel(global.BackGround, timeNow+"_data", field)
				} else {
					delMarshal, err := json.Marshal(dataDetailObj)
					if err != nil {
						return err
					}
					config.RDB.HSet(global.BackGround, timeNow+"_data", field, delMarshal)
				}
				//同步到mysql中
				mysqlItem := model.DataConfirmed{
					Provider:          dataDetailObj.ProvidePerson,
					DataContent:       contentName,
					ContentThumbsUp:   record.ThumbUp,
					ContentThumbsDown: record.ThumbDown,
				}

				mysqlThumbDetail := model.DataThumbDetail{
					ThumbUpPerson:   record.ThumbUpPerson,
					ThumbDownPerson: record.ThumbDownPerson,
				}

				mysqlItem.DataThumbDetail = mysqlThumbDetail
				err = config.DB.Model(&model.Data{}).Create(&mysqlItem).Error
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
