package database

import (
	"mini-project/config"
	"mini-project/models"

	"github.com/labstack/echo"
)

func DashboardAdmin(c echo.Context) (int, int, int, int, int, error) {
	var user models.User
	var adoption models.Adoption
	var pet models.Pet
	var totalUser int
	var totalUserAdopt int
	var totalUserDonate int
	var totalPetAvailable int
	var totalPetAdopted int

	if err := config.DB.Find(&user).Count(&totalUser).Error; err != nil {
		totalUser = 0
	}
	if err := config.DB.Find(&adoption).Count(&totalUserAdopt).Error; err != nil {
		totalUserAdopt = 0
	}
	if err := config.DB.Find(&pet).Count(&totalUserDonate).Error; err != nil {
		totalUserDonate = 0
	}
	if err := config.DB.Where("status = ?", true).Find(&pet).Count(&totalPetAvailable).Error; err != nil {
		totalPetAvailable = 0

	}
	if err := config.DB.Where("status = ?", false).Find(&pet).Count(&totalPetAdopted).Error; err != nil {
		totalPetAdopted = 0
	}
	return totalUser, totalUserAdopt, totalUserDonate, totalPetAvailable, totalPetAdopted, nil
}

func CreatePetCategory(c echo.Context) error {
	var category models.PetCategory
	c.Bind(&category)
	if err := config.DB.Save(&category).Error; err != nil {
		return err
	}
	return nil
}

func GetPetCategory(c echo.Context) (interface{}, error) {
	var category []models.PetCategory
	if err := config.DB.Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func GetPetsAdmin(c echo.Context) (interface{}, error) {
	var pets []models.Pet
	if err := config.DB.Find(&pets).Error; err != nil {
		return nil, err
	}
	return pets, nil
}

func GetUsers(c echo.Context) (interface{}, error) {
	var users []models.User

	if err := config.DB.Preload("UserDetail").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
