package redis

import (
	"cloud-api-go/pkg/logger"
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type RdsClient struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// Redis 全局 Redis
var Redis *RdsClient

// ConnectRedis 连接 redis 并设置全局的 Redis 对象
func ConnectRedis(addr string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(addr, username, password, db)
	})
}

// NewClient 创建一个新的 redis 连接
func NewClient(addr string, username string, password string, db int) *RdsClient {
	// 初始化自定义的 RdsClient 实例
	rds := &RdsClient{}
	// 使用默认的　context
	rds.Context = context.Background()
	// 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       db,
	})
	// 测试连接
	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

// Ping 用于测试 redis 连接是否正常
func (rds *RdsClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储 key 对应的 value 并且设置过期时间 expiration
func (rds *RdsClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func (rds *RdsClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}

	return result
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都会返回 false
func (rds *RdsClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return false
	}

	return true
}
