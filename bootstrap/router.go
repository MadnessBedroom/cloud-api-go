package bootstrap

import (
	"cloud-api-go/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouter(router *gin.Engine) {
	// 注册全局中间件
	// 注册 API 路由
	routes.RegisterApiRoutes(router)
	// 配置 404 路由
	setupNotFoundHandler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
}

func setupNotFoundHandler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		// 获取请求头的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "路由未定义，请确认 url 和请求方法是否正确!",
			})
		}
	})
}
