package common

import (
	"errors"
	"fisco/config"
	"fisco/model"
	"fisco/utils/common"
)

func Login(username, password string) (*model.GeneralUser, error) {
	var user *model.GeneralUser
	config.DB.Model(&model.GeneralUser{}).Where("username = ?", username).Find(&user)
	if user == nil {
		return nil, errors.New("用户已存在，请勿重复创建")
	}
	if user.Username == "" {
		return nil, errors.New("用户不存在，请先注册")
	}
	equal, err := common.Equal(user.Password, password)
	if err != nil {
		return nil, errors.New("密码错误")
	}
	if equal {
		return user, nil
	}
	return nil, errors.New("用户不存在，请先注册")
}

func Role(roleID int) (string, error) {
	var user *model.Role
	err := config.DB.Model(&model.Role{}).Where("role_id = ?", roleID).Find(&user).Error
	if err != nil {
		return "", err
	}
	return user.RoleName, nil
}
