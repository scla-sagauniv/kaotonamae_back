package migrate

import (
	"kaotonamae_back/db"
	"kaotonamae_back/models"
)

func Run() {
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.UserInfo{})
}
