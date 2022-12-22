package controllers

import (
	"cloud-api-go/pkg/auth"
	"cloud-api-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type GamerController struct {
	BaseController
}

// Current 当前登录的选手信息
func (ctrl *GamerController) Current(c *gin.Context) {
	gamerModel := auth.CurrentUser(c)
	response.Data(c, gamerModel)
}

// Index 所有选手
func (ctrl *GamerController) Index(c *gin.Context) {
}
