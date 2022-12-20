package admin

import "cloud-api-go/app/models"

type Admin struct {
	models.BaseModel
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	models.CommonTimestampsField
}

func (a *Admin) Update() {

}

func (a *Admin) Save() {
}
