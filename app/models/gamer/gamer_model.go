package gamer

import "cloud-api-go/app/models"

type Gamer struct {
	models.BaseModel
	Username string
	Password string
	Nickname string
	Avatar   string
	Status   uint64
	models.CommonTimestampsField
}
