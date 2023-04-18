package main

import (
	_ "fisco/config"
	"fisco/load"
	"fisco/routes"
)

// @title     fisco_bcos

// @host      localhost:9091
// @BasePath  /

func main() {
	//client.Client()
	r := routes.NewRouter()

	//r.POST("/login", common.Login)
	//r.POST("/add/role", admin.AddRole)
	//r.POST("/register", common.Register)
	//r.POST("/verify/user", admin.VerifyUser)
	//r.POST("/verify/email", contractuser.Verify)
	//r.POST("/data/provide", datauser.ProvideData)
	//r.POST("/data/delete", datauser.DeleteData)
	//r.GET("/show/hot", generaluser.ShowHotData)
	//r.POST("/data/thumbup", generaluser.ThumbUpData)
	//r.POST("/data/thumbdown", generaluser.ThumbDownData)
	//r.POST("/system/aggregate", system.AggregateToRedis)
	//r.POST("/data/jugde", judge.JudgeUploadData)
	r.Run(":" + load.VP.GetString("server.port"))
}
