package database

import (
	"mini-project/config"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetPets() (interface{}, error) {
	var pets []models.Pet

	if err := config.DB.Preload("PetCategory").Find(&pets).Error; err != nil {
		return nil, err
	}
	return pets, nil
}

func GetPetDetail(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var pet models.Pet
	if err := config.DB.Preload("PetCategory").First(&pet, id).Error; err != nil {
		return nil, err
	}
	return pet, nil
}

func DonatePet(c echo.Context) error {
	var pet models.Pet

	_, id := Authorization(c)
	c.Bind(&pet)
	if err := config.DB.Model(&pet).Create(&models.Pet{
		Deskripsi:     pet.Deskripsi,
		PetCategoryID: pet.PetCategoryID,
		UserID:        id,
	}).Error; err != nil {
		return err
	}
	return nil
}

func AdoptPet(c echo.Context) error {
	var pet models.Pet
	var adoption models.Adoption

	_, user_id := Authorization(c)
	id, _ := strconv.Atoi(c.Param("id"))
	pet_id := uint(id)

	if err := config.DB.Find(&pet, pet_id).Error; err != nil {
		return err
	}

	if pet.UserID == user_id {
		return echo.NewHTTPError(http.StatusBadRequest, "Owner can't adopt the pet")
	}

	if err := config.DB.Model(&pet).Where("id = ?", pet_id).Update(models.Pet{
		Status: true,
	}).Error; err != nil {
		return err
	}
	if err := config.DB.Model(&adoption).Save(&models.Adoption{
		UserID: user_id,
		PetID:  pet_id,
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeletePet(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, user_id := Authorization(c)
	var pet models.Pet
	if err := config.DB.First(&pet, id).Error; err != nil {
		return err
	}

	if pet.UserID != user_id {
		return echo.NewHTTPError(http.StatusBadRequest, "Don't have permission")
	}

	if err := config.DB.Delete(&pet, id).Error; err != nil {
		return err
	}
	return nil
}

// func UpdatePet(c echo.Context) (interface{}, error) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	var pet models.Pet

// 	if err := config.DB.First(&pet, id).Error; err != nil {
// 		return nil, err
// 	}
// 	c.Bind(&pet)
// 	if err := config.DB.Model(&pet).Where("id = ?", id).Updates(models.Pet{
// 		Name:     pet.Name,
// 		Email:    pet.Email,
// 		Password: pet.Password,
// 	}).Error; err != nil {
// 		return nil, err
// 	}
// 	return pet, nil
// }

// func CreatePet(c echo.Context) (interface{}, error) {
// 	pet := models.Pet{}
// 	c.Bind(&pet)

// 	if err := config.DB.Save(&pet).Error; err != nil {
// 		return nil, err
// 	}
// 	return pet, nil
// }

// func CreatePetCategory(c echo.Context) (interface{}, error) {
// 	category := models.PetCategory{}
// 	c.Bind(&category)

// 	if err := config.DB.Save(&category).Error; err != nil {
// 		return nil, err
// 	}
// 	return category, nil
// }
