package game

import (
	"cloud-api-go/app/models"
	"cloud-api-go/pkg/database"
)

type Game struct {
	models.BaseModel
	Title  string
	Desc   string
	Status uint64
	models.CommonTimestampsField
}

func (g *Game) Create() {
	database.DB.Create(&g)
}

func (g *Game) Save() {

}

func (g *Game) Delete() {

}
