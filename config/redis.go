package config

import (
	"fisco/load"
	_ "fisco/load"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	RDB = &redis.Client{}
)

func init() {
	db, _ := strconv.Atoi(load.VP.GetString("redis.db"))
	RDB = redis.NewClient(&redis.Options{
		Addr:     load.VP.GetString("redis.url") + ":" + load.VP.GetString("redis.port"),
		Password: load.VP.GetString("redis.password"),
		DB:       db,
	})
}
