package database

import (
	"mini-project/config"
	"mini-project/models"
	"mini-project/models/payload"

	"github.com/labstack/echo"
)

func DashboardAdmin(c echo.Context) (interface{}, error) {
	var user models.User
	var adoption models.Adoption
	var pet models.Pet
	var dashboard payload.DashboardAdmin

	if err := config.DB.Model(&user).Count(&dashboard.TotalUser).Error; err != nil {
		dashboard.TotalUser = 0
	}
	if err := config.DB.Model(&adoption).Count(&dashboard.TotalAdoption).Error; err != nil {
		dashboard.TotalAdoption = 0
	}
	if err := config.DB.Model(&pet).Count(&dashboard.TotalDonation).Error; err != nil {
		dashboard.TotalDonation = 0
	}
	if err := config.DB.Model(&pet).Where("status = ?", "available").Count(&dashboard.TotalPetAvailable).Error; err != nil {
		dashboard.TotalPetAvailable = 0

	}
	if err := config.DB.Model(&pet).Where("status = ?", "adopted").Count(&dashboard.TotalPetAdopted).Error; err != nil {
		dashboard.TotalPetAdopted = 0
	}
	return dashboard, nil
}

func CreatePetCategory(c echo.Context) error {
	var create payload.CreateCategory
	c.Bind(&create)
	category := models.PetCategory{
		Name: create.Name,
	}
	if err := config.DB.Save(&category).Error; err != nil {
		return err
	}
	return nil
}

func GetPetCategory(c echo.Context) ([]string, error) {
	var category []models.PetCategory
	var response []string
	if err := config.DB.Find(&category).Error; err != nil {
		return nil, err
	}
	for _, value := range category {
		response = append(response, value.Name)
	}
	return response, nil
}

func GetPetsAdmin(c echo.Context) (interface{}, error) {
	var pets []models.Pet
	var response []payload.GetPet
	if err := config.DB.Preload("PetCategory").Find(&pets).Error; err != nil {
		return nil, err
	}
	for _, value := range pets {
		response = append(response, payload.GetPet{
			Deskripsi: value.Deskripsi,
			Status:    value.Status,
			Category:  value.PetCategory.Name,
		})
	}
	return response, nil
}

func GetUsers(c echo.Context) (interface{}, error) {
	var users []models.User
	var response []payload.GetProfil

	if err := config.DB.Preload("UserDetail").Find(&users).Error; err != nil {
		return nil, err
	}

	for _, value := range users {
		response = append(response, payload.GetProfil{
			Name:      value.Name,
			Username:  value.Username,
			Email:     value.Email,
			Alamat:    value.UserDetail.Alamat,
			Handphone: value.UserDetail.Handphone,
		})
	}

	return response, nil
}
