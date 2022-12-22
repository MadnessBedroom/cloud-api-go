package auth

import (
	"cloud-api-go/app/requests"
	"cloud-api-go/pkg/auth"
	"cloud-api-go/pkg/jwt"
	"cloud-api-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

// AdminLogin 管理员登录
func (lc *LoginController) AdminLogin(c *gin.Context) {
	req := requests.AdminLoginRequest{}
	if ok := requests.Validate(c, &req, requests.LoginByAdmin); !ok {
		return
	}

	// 尝试登录
	user, err := auth.Attempt(req.Username, req.Password)
	if err != nil {
		response.Unauthorized(c, "账号不存在或密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Username)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// GamerLogin 选手登录
func (lc *LoginController) GamerLogin(c *gin.Context) {

}
