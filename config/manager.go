package config

import "competition-backend/pkg/config"

func init() {
	config.Add("manager", func() map[string]interface{} {
		return map[string]interface{}{
			"username": config.Env("MANAGER_USERNAME", "admin"),
			"password": config.Env("MANAGER_PASSWORD", "123456"),
		}
	})
}
