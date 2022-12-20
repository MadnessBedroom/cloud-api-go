package gamer

import "cloud-api-go/pkg/database"

// Get 通过 ID 获取选手
func Get(idStr string) (gamerModel Gamer) {
	database.DB.Where("id", idStr).First(&gamerModel)
	return
}

// All 获取所有选手
func All() (gamers []Gamer) {
	database.DB.Find(&gamers)
	return
}

// IsUsernameExist 判断选手用户名是否存在
func IsUsernameExist(username string) bool {
	var count int64
	database.DB.Model(Gamer{}).Where("username = ?", username).Count(&count)
	return count > 0
}
