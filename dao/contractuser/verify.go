package contractuser

import (
	"errors"
	"fisco/config"
	"fisco/model"
	"fisco/utils/common"
	"fmt"
)

func Verify(username, verifyEmail string) error {
	user := model.ContractUser{}
	config.DB.Model(&model.ContractUser{}).Where("username=?", username).Find(&user)
	if user.VerifyEmail != verifyEmail {
		return errors.New("申请失败，邮件错误")
	} else {
		token, _ := common.ParseToken(verifyEmail)
		err := token.Valid()
		if err != nil {
			return errors.New("申请失败，邮件过期")
		} else {
			fmt.Println(user.Verified)
			if user.Verified == true {
				return errors.New("您已确认过邮件，请勿重复确认")
			}
		}
	}
	err := config.DB.Model(&model.ContractUser{}).Where("username=?", username).Update("verified", 1).Error
	if err != nil {
		return err
	}
	err = config.DB.Model(&model.GeneralUser{}).Where("username=?", username).Update("state", 1).Error
	return err
}
