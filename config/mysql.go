package config

import (
	"fisco/load"
	_ "fisco/load"
	"fisco/model"

	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	DSN string
)

func init() {
	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		load.VP.GetString("mysql.user"),
		load.VP.GetString("mysql.password"),
		load.VP.GetString("mysql.url"),
		load.VP.GetString("mysql.port"),
		load.VP.GetString("mysql.dbname"))
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&model.Role{},
		&model.GeneralUser{},
		&model.ContractUser{},
		&model.Verify{},
		&model.Data{},
		&model.DataProvider{},
		&model.DataConfirmed{},
		&model.DataWatcher{},
		&model.DataJudge{})
}
