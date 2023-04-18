package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type DataConfirmed struct {
	gorm.Model
	Provider          string          `json:"data_provider" gorm:"data_provider"`
	DataName          string          `json:"data_name" gorm:"data_name"`
	DataContent       string          `json:"data_content" gorm:"data_content"`
	ContentThumbsUp   int             `json:"data_thumbs_up" gorm:"data_thumbs_up"`
	ContentThumbsDown int             `json:"data_thumbs_down" gorm:"data_thumbs_down"`
	DataThumbDetail   DataThumbDetail `json:"data_thumb_detail" gorm:"data_thumb_detail"`
}

type DataThumbDetail struct {
	ThumbUpPerson   []string `json:"thumb_up_person" gorm:"thumb_up_person"`
	ThumbDownPerson []string `json:"thumb_down_person" gorm:"thumb_down_person"`
}

func (d DataThumbDetail) Value() (driver.Value, error) {
	str, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return string(str), nil
}

func (d *DataThumbDetail) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("类型不匹配")
	}
	err := json.Unmarshal(bytes, d)
	if err != nil {
		return err
	}
	return nil
}
