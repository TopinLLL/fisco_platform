package admin

import (
	"errors"
	"fisco/config"
	"fisco/load"
	"fisco/model"
	"fisco/utils/common"
)

func VerifyUser(userID int, ok bool) (*model.Verify, error) {
	existed := config.DB.Model(&model.Verify{}).Limit(1).Where("id=?", userID).RowsAffected
	if existed != 0 {
		return nil, errors.New("此用户不存在")
	}

	var verifyUser model.Verify
	config.DB.Model(&model.Verify{}).Where("id = ?", userID).Find(&verifyUser)

	//不成功发送邮件
	if ok == false {
		config.DB.Model(&model.Verify{}).Delete(&model.Verify{}, userID)
		common.SendMail(load.VP.GetString("mail.from"), verifyUser.Mail, "对不起您的账户未申请成功", "详情请联系管理员")
		return nil, nil
	}

	config.DB.Model(&model.Verify{}).Delete(&model.Verify{}, userID)

	//如果是智能合约用户，则需要生成密钥再插入数据库
	if verifyUser.RoleId == 2 {
		pubKey, priKey, address, err := common.GenerateKey()
		if err != nil {
			return nil, errors.New("智能合约用户创建失败，请联系管理员")
		} else {
			contractUser := model.ContractUser{
				Username:       verifyUser.Username,
				PrivateKey:     priKey,
				PublicKey:      pubKey,
				Address:        address,
				ContractNumber: 0,
				ViolationCount: 0,
			}

			//构造邮件token
			tokenString, _ := common.GenerateToken(contractUser.UserID, contractUser.Username)
			contractUser.VerifyEmail = tokenString
			common.SendMail(load.VP.GetString("mail.from"), verifyUser.Mail, "您的账户申请成功", "请使用以下token进行验证"+tokenString)
			config.DB.Model(&model.ContractUser{}).Create(&contractUser)
		}
	}
	config.DB.Model(&model.GeneralUser{}).Where("username=?", verifyUser.Username).Update("state", 1)
	return &verifyUser, nil
}
