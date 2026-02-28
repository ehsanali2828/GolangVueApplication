package initializers

import "github.com/robyklein/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
