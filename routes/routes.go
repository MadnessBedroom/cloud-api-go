package routes

import (
	"cloud-api-go/app/controllers/auth"
	"cloud-api-go/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(router *gin.Engine) {
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = router.Group("/api/v1")
	} else {
		v1 = router.Group("/v1")
	}

	{
		// http://localhost:3000/api/v1/auth/game/list
		authGroup := v1.Group("/auth")
		{
			// 管理员和选手登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/admin", lgc.AdminLogin)
			authGroup.POST("/login/gamer", lgc.GamerLogin)

			// 比赛的增删改查
			authGroup.POST("/game/add")
			authGroup.GET("/game/list")
			authGroup.PATCH("/game/:id")
			authGroup.DELETE("/game/:id")
		}
	}
}
