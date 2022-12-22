package admin

import "cloud-api-go/pkg/database"

// GetByUsername 通过用户名来获取用户信息
func GetByUsername(username string) (adminModel Admin) {
	database.DB.Where("username = ?", username).First(&adminModel)
	return
}
