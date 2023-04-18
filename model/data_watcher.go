package model

import "gorm.io/gorm"

type DataWatcher struct {
	gorm.Model
	Username           string `json:"username" gorm:"username"`
	ThumbUpContentID   int    `json:"thumb_up_content_id" gorm:"thumb_up_content_id"`
	ThumbDownContentID int    `json:"thumb_down_content_id" gorm:"thumb_down_content_id"`
	//信誉分
	Reputation float64 `json:"reputation" gorm:"reputation"`
}
