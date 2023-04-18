package model

import (
	"gorm.io/gorm"
)

type DataJudge struct {
	gorm.Model
	DataProvider string `json:"data_provider" gorm:"data_provider"`
	DataContent  string `json:"data_content" gorm:"data_content"`
	JudgeResult  string `json:"judge_result" gorm:"judge_result"`
}

//type Provider struct {
//	ID   int    `json:"id" gorm:"id"`
//	Name string `json:"name" gorm:"name"`
//}
//
//func (p Provider) Value() (driver.Value, error) {
//	marshal, err := json.Marshal(p)
//	if err != nil {
//		return nil, err
//	}
//	return string(marshal), nil
//}
//
//func (p *Provider) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return errors.New("类型不匹配")
//	}
//	err := json.Unmarshal(bytes, p)
//	if err != nil {
//		return err
//	}
//	return nil
//}
