package faker

import (
	"mini-project/models"

	"github.com/jinzhu/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		Name:     "admin",
		Username: "admin",
		Password: "admin",
		Role:     "admin",
	}
}
