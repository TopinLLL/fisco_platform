package model

import "gorm.io/gorm"

type GeneralUser struct {
	gorm.Model
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	RoleID   int    `json:"role_id" gorm:"role_id"`
	State    bool   `json:"state" gorm:"state"`
	Mail     string `json:"mail" gorm:"mail"`
}
