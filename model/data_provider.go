package model

type DataProvider struct {
	UserID           int    `json:"user_id" gorm:"user_id"`
	UserName         string `json:"user_name" gorm:"user_name"`
	ProvideContentID int    `json:"provide_data_id" gorm:"provide_data_id"`
	//信誉分
	Reputation float64 `json:"reputation" gorm:"reputation"`
	//被举报次数
	UserThumbDown int `json:"user_thumb_down" gorm:"user_thumb_down"`
}
