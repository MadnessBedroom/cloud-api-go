package gamer

import (
	"cloud-api-go/app/models"
	"cloud-api-go/pkg/database"
	"cloud-api-go/pkg/hash"
)

type Gamer struct {
	models.BaseModel
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Status   uint64 `json:"status,omitempty"`
	models.CommonTimestampsField
}

// Create 创建用户
func (g *Gamer) Create() {
	database.DB.Create(&g)
}

// ComparePassword 检查密码是否正确
func (g *Gamer) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, g.Password)
}

func (g *Gamer) Save() (rowsAffected int64) {
	result := database.DB.Save(&g)
	return result.RowsAffected
}
