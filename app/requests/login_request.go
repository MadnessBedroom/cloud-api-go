package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AdminLoginRequest struct {
	Username      string `json:"username,omitempty" valid:"username"`
	Password      string `json:"password,omitempty" valid:"password"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
}

// LoginByAdmin 管理员登录的验证表单，errs 的长度等于零表示验证通过
func LoginByAdmin(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"username":       []string{"required", "min:4"},
		"password":       []string{"required", "min:6"},
		"captcha_answer": []string{"required", "digits:5"},
	}

	messages := govalidator.MapData{
		"username": []string{
			"required:用户名为必填项，参数名称 username",
			"min:用户名长度需大于等于 5",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于等于 6",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 5 位",
		},
	}
	errs := validate(data, rules, messages)

	_data := data.(*AdminLoginRequest)

	return errs
}

// LoginByGamer 选手登录的验证表单，errs 的长度等于零表示验证通过
func LoginByGamer(data interface{}, c *gin.Context) map[string][]string {

}
