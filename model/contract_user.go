package model

type ContractUser struct {
	UserID         int    `json:"user_id" gorm:"user_id" `
	Username       string `json:"username" gorm:"username"`
	PrivateKey     string `json:"private_key" gorm:"private_key"`
	PublicKey      string `json:"public_key" gorm:"public_key"`
	Address        string `json:"address" gorm:"address"`
	ContractNumber int    `json:"contract_number" gorm:"contract_number"`
	ViolationCount int    `json:"violation_count" gorm:"violation_count"`
	Verified       bool   `json:"verified" gorm:"verified"`
	VerifyEmail    string `json:"verify_email" gorm:"verify_email"`
	//信誉分
	Reputation float64 `json:"reputation" gorm:"reputation"`
	//被举报次数
	UserThumbDown int `json:"user_thumb_down" gorm:"user_thumb_down"`
}
