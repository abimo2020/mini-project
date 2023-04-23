package database

import (
	"mini-project/config"
	"mini-project/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Preload("UserDetail").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := config.DB.Preload("UserDetail").First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	c.Bind(&user)
	if err := config.DB.Model(&user).Where("id = ?", id).Updates(models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}).Error; err != nil {
		return nil, err
	}
	return user, nil
}
