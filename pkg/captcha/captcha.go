package captcha

import (
	"cloud-api-go/pkg/app"
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

// GenerateCaptcha 生成图片验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证验证码是否正确
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {

	// 方便本地和 API 自动测试
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}
	// 第三个参数是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单提交错误需要多次输入图片验证码
	return c.Base64Captcha.Verify(id, answer, false)
}
