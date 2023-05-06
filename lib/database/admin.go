package database

import (
	"mini-project/config"
	"mini-project/models"
)

func CreatePetCategory(req *models.PetCategory) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func GetPetCategory() (category []models.PetCategory, err error) {
	if err := config.DB.Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
