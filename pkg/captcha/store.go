package captcha

import (
	"cloud-api-go/pkg/app"
	"cloud-api-go/pkg/config"
	"cloud-api-go/pkg/redis"
	"errors"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RdsClient
	KeyPrefix   string
}

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (rs *RedisStore) Set(key string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	// 方便本地开发调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}

	if ok := rs.RedisClient.Set(rs.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}

	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (rs *RedisStore) Get(key string, clear bool) string {
	newKey := rs.KeyPrefix + key
	val := rs.RedisClient.Get(newKey)
	if clear {
		rs.RedisClient.Del(newKey)
	}

	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法证码的答案
func (rs *RedisStore) Verify(key, answer string, clear bool) bool {
	v := rs.Get(key, clear)
	return v == answer
}
