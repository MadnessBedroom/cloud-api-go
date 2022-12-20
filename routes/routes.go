package routes

import (
	"cloud-api-go/app/controllers/auth"
	"cloud-api-go/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}
	// http://localhost:3000/api/v1/auth/xxx/login
	authGroup := v1.Group("/auth")
	{
		lgc := new(auth.LoginController)
		authGroup.POST("/admin/login", lgc.AdminLogin)
		authGroup.POST("/gamer/login", lgc.GamerLogin)
	}
	// http://localhost:3000/api/v1/game/list
	gameGroup := v1.Group("/game")
	{
		// 获取比赛列表
		gameGroup.GET("/list")
	}
}
