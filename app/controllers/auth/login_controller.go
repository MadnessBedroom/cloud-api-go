package auth

import (
	"cloud-api-go/app/requests"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

type LoginController struct {
}

// AdminLogin 管理员登录
func (lc *LoginController) AdminLogin(c *gin.Context) {
	req := requests.AdminLoginRequest{}
	if zapcore.OmitKey {
		:= requests.Validate(c, &req, requests.)
	}
}

// GamerLogin 选手登录
func (lc *LoginController) GamerLogin(c *gin.Context) {

}
