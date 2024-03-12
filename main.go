package main

import (
	_ "fisco/config"
	"fisco/load"
	"fisco/routes"
)

// @title     面向电子合同的可修正区块链系统
// @host      localhost:9091
// @BasePath  /
func main() {
	r := routes.NewRouter()
	r.Run(":" + load.VP.GetString("server.port"))
}
