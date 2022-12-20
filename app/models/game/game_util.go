package game

import "cloud-api-go/pkg/database"

func IsGameTitleExist(title string) bool {
	var count int64
	database.DB.Model(Game{}).Where("title = ?", title).Count(&count)
	return count > 0
}
