package captcha

import (
	"cloud-api-go/pkg/config"
	"cloud-api-go/pkg/redis"
	"github.com/mojocn/base64Captcha"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// once 确保 internalCaptcha 对象只初始化一次
var once sync.Once

var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	once.Do(func() {
		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}
		// redis
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha:",
		}
		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),
			config.GetInt("captcha,width"),
			config.GetInt("captcha.length"),
			config.GetFloat64("captcha.max_skew"),
			config.GetInt("captcha.dot_count"),
		)

		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha
}
