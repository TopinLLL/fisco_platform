package model

import "gorm.io/gorm"

type Verify struct {
	gorm.Model
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	Mail     string `json:"mail" gorm:"mail"`
	RoleId   int    `json:"role_id" gorm:"role_id"`
}
