package middlewares

import (
	"cloud-api-go/app/models/gamer"
	"cloud-api-go/pkg/config"
	"cloud-api-go/pkg/jwt"
	"cloud-api-go/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(c)
		// 解析失败
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}
		gamerModel := gamer.Get(claims.UserID)
		if gamerModel.ID == 0 {
			response.Unauthorized(c, "找不到相应用户，用户可能已删除")
			return
		}

		c.Set("current_user_id", gamerModel.GetStringID())
		c.Set("current_user_name", gamerModel.Username)
		c.Set("current_user", gamerModel)
		c.Next()
	}
}
