package model

type Role struct {
	RoleID int `json:"role_id" gorm:"role_id"`
	RoleName string `json:"role_name" gorm:"role_name"`
}
