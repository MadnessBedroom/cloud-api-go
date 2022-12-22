package auth

import (
	"cloud-api-go/app/models/admin"
	"cloud-api-go/app/models/gamer"
	"cloud-api-go/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
)

// Attempt 尝试登录
func Attempt(username, password string) (admin.Admin, error) {
	adminModel := admin.GetByUsername(username)
	if adminModel.ID == 0 {
		return admin.Admin{}, errors.New("账号不存在")
	}
	if !adminModel.ComparePassword(password) {
		return admin.Admin{}, errors.New("密码错误")
	}

	return adminModel, nil
}

// CurrentUser 从 gin.context 中获取当前的登录用户
func CurrentUser(c *gin.Context) gamer.Gamer {
	gamerModel, ok := c.MustGet("current_user").(gamer.Gamer)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return gamer.Gamer{}
	}

	return gamerModel
}
