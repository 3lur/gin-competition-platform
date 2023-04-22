package main

import (
	"competition-backend/bootstrap"
	btsConfig "competition-backend/config"
	configPkg "competition-backend/pkg/config"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	// 初始化配置
	configPkg.SetupConfig("")

	// 初始化日志
	bootstrap.SetupLogger()

	// 初始化数据库
	bootstrap.SetupDB()

	// 初始化 Redis
	bootstrap.SetupRedis()

	// 初始化管理员信息
	bootstrap.SetupManager()

	// 初始化路由
	bootstrap.SetupRouter()
}
