package main

import (
	"softwareCup/router"
	"softwareCup/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMySQL()
	utils.Init(1)
	r := router.Router()

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}
