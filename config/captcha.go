package config

import "cloud-api-go/pkg/config"

func init() {
	config.Add("captcha", func() map[string]interface{} {
		return map[string]interface{}{
			"height":            80,                  // 验证码图片高度
			"width":             240,                 // 验证码图片宽度
			"length":            5,                   // 验证码答案长度
			"max_skew":          0.7,                 // 数字的最大倾斜度
			"dot_count":         80,                  // 图片背景里的混淆点数量
			"expire_time":       15,                  // 过期时间，单位分钟
			"debug_expire_time": 10080,               // debug 模式下的过期时间
			"testing_key":       "captcha_skip_test", // 非 production 环境下，使用此 key 跳过验证，方便测试
		}
	})
}
