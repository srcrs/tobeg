package main

import (
	"fmt"
	"tobeg/global"
	"tobeg/initialize"
)

func main() {
	// 初始化日志
	initialize.InitLogger()

	// 初始化配置信息
	initialize.InitConfig()
	// 初始化证书配置
	initialize.InitCert()

	// 初始化数据库
	initialize.InitSQLite3()

	// 初始化路由
	Router := initialize.Routers()

	startErr := Router.Run(fmt.Sprintf(":%d", global.Config.ServerConfig.Port))

	if startErr != nil {
		fmt.Println("启动失败")
	}
}
