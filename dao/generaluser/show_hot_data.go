package generaluser

import (
	"encoding/json"
	"errors"
	"fisco/config"
	"fisco/dto"
	"fisco/global"
	"time"

	"github.com/go-redis/redis/v8"
)

func ShowHotData() ([]redis.Z, error) {
	var err error
	nowTime := time.Now().Format("2006-01-02")
	//当日热点数据
	key, err := config.RDB.Keys(global.BackGround, nowTime+"_data").Result()
	if err != nil {
		return nil, err
	}
	if len(key) == 0 {
		return nil, errors.New("暂时没有数据")
	} else {
		keys, err := config.RDB.HKeys(global.BackGround, nowTime+"_data").Result()
		if err != nil {
			return nil, err
		}
		for _, item := range keys {
			value, err := config.RDB.HGet(global.BackGround, nowTime+"_data", item).Result()
			if err != nil {
				return nil, err
			}

			today := &dto.EveryDayData{}
			err = json.Unmarshal([]byte(value), today)
			if err != nil {
				return nil, err
			}
			//热点数据ZSET重排
			for contentName, detail := range today.DataContent {
				err = config.RDB.ZAdd(global.BackGround, nowTime+"_hot_data", &redis.Z{
					Score:  float64(detail.ThumbUp) - float64(detail.ThumbDown),
					Member: contentName,
				}).Err()
			}
			config.RDB.Expire(global.BackGround, nowTime+"_hot_data", time.Hour*24)
			if err != nil {
				return nil, err
			}
		}

	}

	result, err := config.RDB.ZRangeWithScores(global.BackGround, nowTime+"_hot_data", 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
