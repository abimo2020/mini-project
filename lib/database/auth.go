package database

import (
	"mini-project/config"
	"mini-project/middlewares"
	"mini-project/models"

	"github.com/labstack/echo"
)

func LoginUser(user *models.User) (interface{}, error) {
	var err error
	var token models.Token
	if err = config.DB.Where("email = ? OR username = ?  AND password = ? ", user.Email, user.Username, user.Password).First(&user).Error; err != nil {
		return nil, err
	}
	token.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func Register(c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}