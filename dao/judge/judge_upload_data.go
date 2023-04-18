package judge

import (
	"fisco/config"
	"fisco/model"
)

func JudgeUploadData(agree bool, dataID int) error {
	data := model.Data{}
	err := config.DB.Model(&model.Data{}).Where("id=?", dataID).Find(&data).Error
	if err != nil {
		return err
	}
	if agree {
		dataJudge := model.DataJudge{
			DataProvider: data.Provider,
			DataContent:  data.DataContent,
			JudgeResult:  "数据已成功上传至主平台",
		}
		dataConfirmed := model.DataConfirmed{
			Provider:          data.Provider,
			DataName:          data.DataName,
			DataContent:       data.DataContent,
			ContentThumbsUp:   data.ContentThumbsUp,
			ContentThumbsDown: data.ContentThumbsDown,
			DataThumbDetail:   data.DataThumbDetail,
		}
		err = config.DB.Model(&model.DataJudge{}).Create(&dataJudge).Error
		if err != nil {
			return err
		}
		err = config.DB.Model(&model.DataConfirmed{}).Create(&dataConfirmed).Error
		if err != nil {
			return err
		}
	} else {
		dataJudge := model.DataJudge{
			DataProvider: data.Provider,
			DataContent:  data.DataContent,
			JudgeResult:  "数据未被允许通过",
		}
		err = config.DB.Model(&model.DataJudge{}).Create(&dataJudge).Error
		if err != nil {
			return err
		}
	}
	//删除临时后台数据
	err = config.DB.Model(&model.Data{}).Where("id=?", dataID).Delete(nil).Error
	if err != nil {
		return err
	}
	return nil
}
