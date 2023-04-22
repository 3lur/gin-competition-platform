package bootstrap

import (
	"competition-backend/app/http/models/manager"
	"competition-backend/pkg/config"
	"competition-backend/pkg/database"
	"competition-backend/pkg/logger"
	"fmt"
)

func SetupManager() {
	var count int64
	database.DB.Table("manager").Count(&count)
	if count > 0 {
		return
	}
	err := createManager()
	if err != nil {
		logger.ErrorString("Bootstrap", "Manager", err.Error())
		fmt.Printf("初始化管理员失败: %v", err.Error())
		return
	}
}

func createManager() error {
	result := database.DB.Table("manager").Create(&manager.Manager{
		Username: config.Get("manager.username"),
		Password: config.Get("manager.password"),
	})
	return result.Error
}
