package bootstrap

import (
	"cloud-api-go/pkg/config"
	"cloud-api-go/pkg/redis"
	"fmt"
)

// SetupRedis 初始化 Redis 连接信息
func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
