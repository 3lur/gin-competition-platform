package bootstrap

import (
	"competition-backend/pkg/config"
	"competition-backend/pkg/redis"
	"fmt"
)

func SetupRedis() {
	redis.Conn(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
