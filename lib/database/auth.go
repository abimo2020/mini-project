package database

import (
	"mini-project/config"

	"mini-project/models"
	"mini-project/models/payload"
)

func Login(req *payload.LoginForm) (user models.User, err error) {
	user = models.User{}

	if err = config.DB.Where("username = ? OR email = ? AND password = ?", req.ID, req.ID, req.Password).First(&user).Error; err != nil {
		return user, err
	}

	return
}

func Register(user *models.User) error {

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
