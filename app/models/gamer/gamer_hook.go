package gamer

import (
	"cloud-api-go/pkg/hash"
	"gorm.io/gorm"
)

// BeforeSave Gorm 的模型钩子，在创建和更新模型前调用
func (g *Gamer) BeforeSave(db *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(g.Password) {
		g.Password = hash.BcryptHash(g.Password)
	}

	return
}
