package main

import (
	_ "fisco/config"
	"fisco/load"
	"fisco/routes"
)

// @title     区块链数据可修改实验平台
// @host      localhost:9091
// @BasePath  /
func main() {
	r := routes.NewRouter()
	r.Run(":" + load.VP.GetString("server.port"))
}
