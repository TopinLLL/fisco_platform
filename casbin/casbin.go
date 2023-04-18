package casbin

import (
	"fisco/config"
	"fisco/load"
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	E   *casbin.Enforcer
	err error
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		load.VP.GetString("mysql.user"),
		load.VP.GetString("mysql.password"),
		load.VP.GetString("mysql.url"),
		load.VP.GetString("mysql.port"))
	A, err := gormadapter.NewAdapter("mysql", dsn)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	E, err = casbin.NewEnforcer("casbin/model/user_model.conf", A)
	if err != nil {
		config.Logger.Error(err.Error())
	}
}
