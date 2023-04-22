package cache

import (
	"competition-backend/pkg/config"
	"competition-backend/pkg/redis"
)

type RedisStore struct {
	RedisClient *redis.RdsClient
	KeyPrefix   string
}

func NewRedisStore(addr, username, password string, db int) *RedisStore {
	rs := &RedisStore{}
	rs.RedisClient = redis.NewClient(addr, username, password, db)
	rs.KeyPrefix = config.GetString("app.name") + ":cache:"
	return rs
}
