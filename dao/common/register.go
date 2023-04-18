package common

import (
	"errors"
	"fisco/config"
	"fisco/model"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, username string, password string, roleID int, mail string) (*model.Verify, error) {
	var verify model.Verify
	userFind := config.DB.Model(&model.GeneralUser{}).Where("username=?", username).Find(nil).RowsAffected
	if userFind != 0 {
		return nil, errors.New("用户已经存在，请勿重复申请")
	}
	existed := config.DB.Model(&model.Verify{}).Limit(1).Where("username=?", username).Find(&verify).RowsAffected
	if existed != 0 {
		return nil, errors.New("您已经申请过了,请勿重复提交")
	}

	//verify
	verify = model.Verify{
		Username: username,
		Password: password,
		Mail:     mail,
		RoleId:   roleID,
	}
	err := config.DB.Model(&model.Verify{}).Create(&verify).Error

	//注册用户
	user := &model.GeneralUser{
		Username: verify.Username,
		Password: verify.Password,
		RoleID:   verify.RoleId,
		Mail:     verify.Mail,
	}

	config.DB.Model(&model.GeneralUser{}).Create(user)

	switch roleID {
	//数据观测用户
	case 3:
		dataWatcher := model.DataWatcher{
			Username:           username,
			ThumbUpContentID:   0,
			ThumbDownContentID: 0,
			Reputation:         0,
		}
		err = config.DB.Model(&model.DataWatcher{}).Create(&dataWatcher).Error
		if err != nil {
			return nil, err
		}
	//数据提供用户
	case 4:
		user := model.GeneralUser{}
		err = config.DB.Model(&model.GeneralUser{}).Where("username=?", username).Find(&user).Error
		if err != nil {
			return nil, err
		}
		dataProvider := model.DataProvider{
			UserID:           user.RoleID,
			UserName:         username,
			ProvideContentID: 0,
			Reputation:       0,
			UserThumbDown:    0,
		}
		err = config.DB.Model(&model.DataProvider{}).Create(&dataProvider).Error
		if err != nil {
			return nil, err
		}
	}

	//先尝试获取是否有该session
	session, err := config.Store.Get(ctx.Request, "sessionID")
	if err != nil {
		return nil, err
	}
	//防止有session的情况重复登录
	session.Options.MaxAge = -1
	//创建新的session
	session, err = config.Store.New(ctx.Request, "sessionID")
	//给session添加信息
	session.Values["username"] = username
	session.Options.MaxAge = 86400
	if err != nil {
		return nil, err
	}

	//将session存入redis
	err = config.Store.Save(ctx.Request, ctx.Writer, session)
	if err != nil {
		return nil, err
	}
	return &verify, nil
}
