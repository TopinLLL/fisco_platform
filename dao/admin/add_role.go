package admin

import (
	"errors"
	"fisco/config"
	"fisco/model"
)

func AddRole(roleID int, roleName string) (*model.Role, error) {
	role := model.Role{
		RoleID:   roleID,
		RoleName: roleName,
	}
	err := config.DB.Model(&model.Role{}).Create(&role).Error
	if err != nil {
		return nil, errors.New("增加用户种类错误")
	}
	return &role, nil
}
