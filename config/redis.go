package config

import "competition-backend/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":           config.Env("REDIS_HOST", "127.0.0.1"),
			"port":           config.Env("REDIS_PORT", "6379"),
			"password":       config.Env("REDIS_PASSWORD", ""),
			"database":       config.Env("REDIS_MAIN_DB", 0),
			"database_cache": config.Env("REDIS_CACHE_DB", 0),
		}
	})
}
