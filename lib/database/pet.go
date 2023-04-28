package database

import (
	"mini-project/config"
	"mini-project/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetPets() (interface{}, error) {
	var pets []models.Pet

	if err := config.DB.Find(&pets).Error; err != nil {
		return nil, err
	}
	return pets, nil
}

func GetPetDetail(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var pet models.Pet
	if err := config.DB.First(&pet, id).Error; err != nil {
		return nil, err
	}
	return pet, nil
}

func DonatePet() {

}

func AdoptPet() {

}

func DeletePet(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var pet models.Pet
	if err := config.DB.First(&pet, id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&pet, id).Error; err != nil {
		return nil, err
	}
	return pet, nil
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

func CreatePet(c echo.Context) (interface{}, error) {
	pet := models.Pet{}
	c.Bind(&pet)

	if err := config.DB.Save(&pet).Error; err != nil {
		return nil, err
	}
	return pet, nil
}

func CreatePetCategory(c echo.Context) (interface{}, error) {
	category := models.PetCategory{}
	c.Bind(&category)

	if err := config.DB.Save(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
