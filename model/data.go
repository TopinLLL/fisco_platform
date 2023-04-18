package model

import (
	"gorm.io/gorm"
)

type Data struct {
	gorm.Model
	Provider          string          `json:"data_provider" gorm:"data_provider"`
	DataName          string          `json:"data_name" gorm:"data_name"`
	DataContent       string          `json:"data_content" gorm:"data_content"`
	ContentThumbsUp   int             `json:"data_thumbs_up" gorm:"data_thumbs_up"`
	ContentThumbsDown int             `json:"data_thumbs_down" gorm:"data_thumbs_down"`
	DataThumbDetail   DataThumbDetail `json:"data_thumb_detail" gorm:"data_thumb_detail"`
}
