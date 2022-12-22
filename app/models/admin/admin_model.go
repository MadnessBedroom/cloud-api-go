package admin

import (
	"cloud-api-go/app/models"
	"cloud-api-go/pkg/hash"
)

type Admin struct {
	models.BaseModel
	Username string `gorm:"type:varchar(50);not null;unique;comment:'管理员登录用户名'" json:"username"`
	Password string `gorm:"size:255;not null;comment:'管理员密码'" json:"password"`
	Nickname string `gorm:"type:varchar(50);comment:'昵称'" json:"nickname"`
	Avatar   string `gorm:"type:varchar(255);comment:'头像'" json:"avatar"`
	models.CommonTimestampsField
}

func (a *Admin) Update() {

}

func (a *Admin) Save() {
}

func (a *Admin) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, a.Password)
}
