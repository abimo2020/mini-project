package database

import (
	"mini-project/config"
	"mini-project/models"
	"mini-project/models/payload"
)

func GetAvailablePets() (pets []models.Pet, err error) {
	if err := config.DB.Preload("PetCategory").Where("status = ?", "available").Find(&pets).Error; err != nil {
		return []models.Pet{}, err
	}
	return pets, nil
}

func GetPets() (pets []models.Pet, err error) {
	if err := config.DB.Preload("PetCategory").Find(&pets).Error; err != nil {
		return []models.Pet{}, err
	}
	return pets, nil
}

func GetPet(id uint) (pet models.Pet, err error) {
	if err := config.DB.Preload("PetCategory").Preload("User.UserDetail").First(&pet, id).Error; err != nil {
		return models.Pet{}, err
	}

	return pet, nil
}

func GetDonateList(id uint) (user models.User, err error) {
	if err := config.DB.Preload("Pet.PetCategory").First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetAdoptList(id uint) (user models.User, err error) {
	if err := config.DB.Preload("Adoption.Pet.PetCategory").Preload("Adoption.Pet.User.UserDetail").First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdatePet(req *payload.UpdatePet, id uint) error {
	var pet models.Pet
	if err := config.DB.Model(&pet).Where("id = ?", id).Update(models.Pet{
		Deskripsi:     req.Deskripsi,
		PetCategoryID: req.PetCategoryID,
	}).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePetStatus(id uint, pet *models.Pet) error {
	if pet.Status == "available" {
		if err := config.DB.Model(&pet).Where("id = ?", id).Update(models.Pet{
			Status: "adopted",
		}).Error; err != nil {
			return err
		}
	} else {
		if err := config.DB.Model(&pet).Where("id = ?", id).Update(models.Pet{
			Status: "available",
		}).Error; err != nil {
			return err
		}
	}
	return nil
}
func AdoptedStatus(id uint) error {
	var pet models.Pet
	if err := config.DB.Model(&pet).Where("id = ? AND status = ?", id, "available").Update(models.Pet{
		Status: "adopted",
	}).Error; err != nil {
		return err
	}
	return nil
}

func DonatePet(req *payload.CreatePet, id uint) error {
	var pet models.Pet

	if err := config.DB.Model(&pet).Create(&models.Pet{
		Deskripsi:     req.Deskripsi,
		PetCategoryID: req.PetCategoryID,
		UserID:        id,
	}).Error; err != nil {
		return err
	}
	return nil
}

func AdoptPet(id uint, petId uint) error {
	var adoption models.Adoption
	if err := config.DB.Model(&adoption).Save(&models.Adoption{
		UserID: id,
		PetID:  petId,
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeletePet(id uint) error {
	var pet models.Pet
	if err := config.DB.Delete(&pet, id).Error; err != nil {
		return err
	}
	return nil
}

func CountPets() (result int) {
	var pet models.Pet
	if err := config.DB.Model(&pet).Count(&result).Error; err != nil {
		return 0
	}
	return result
}

func CountAvailablePets() (result int) {
	var pet models.Pet
	if err := config.DB.Model(&pet).Where("status = ?", "available").Count(&result).Error; err != nil {
		return 0
	}
	return result
}

func CountAdoptedPets() (result int) {
	var pet models.Pet
	if err := config.DB.Model(&pet).Where("status = ?", "adopted").Count(&result).Error; err != nil {
		return 0
	}
	return result
}

func CountAdoption() (result int) {
	var adoption models.Adoption
	if err := config.DB.Model(&adoption).Count(&result).Error; err != nil {
		return 0
	}
	return result
}
