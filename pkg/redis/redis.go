package redis

import (
	"competition-backend/pkg/logger"
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
)

type RdsClient struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// Redis 全局 Redis
var Redis *RdsClient

// Conn 连接 redis 数据库，设置全局的 Redis 对象
func Conn(addr, uname, pwd string, db int) {
	once.Do(func() {
		Redis = NewClient(addr, uname, pwd, db)
	})
}

// NewClient 创建一个新的 redis 连接
func NewClient(addr, uname, pwd string, db int) *RdsClient {
	// 初始化 redis 实例
	rds := &RdsClient{}
	// 使用默认的 context
	rds.Context = context.Background()
	// 使用 redis 库初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: uname,
		Password: pwd,
		DB:       db,
	})
	// 测试连接
	err := rds.Ping()
	logger.LogIf(err)
	return rds
}

// Ping 测试连接是否正常
func (r RdsClient) Ping() error {
	_, err := r.Client.Ping(r.Context).Result()
	return err
}

func (r RdsClient) Set(key string, value interface{}) bool {
	if err := r.Client.Set(r.Context, key, value, 0).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

func (r RdsClient) Get(key string) string {
	res, err := r.Client.Get(r.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return res
}

func (r RdsClient) Has(key string) bool {
	_, err := r.Client.Get(r.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

func (r RdsClient) Del(keys ...string) bool {
	if err := r.Client.Del(r.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}
