package config

import (
	"gopkg.in/boj/redistore.v1"
)

var (
	Store *redistore.RediStore
	err   error
)

func init() {
	Store, err = redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	Store.Options.Secure = false
	Store.Options.HttpOnly = false
	if err != nil {
		Logger.Error(err.Error())
	}
	//defer Store.Close()
}
