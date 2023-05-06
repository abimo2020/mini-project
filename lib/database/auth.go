package database

import (
	"mini-project/config"

	"mini-project/models"
)

func Register(user *models.User) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
