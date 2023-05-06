package database

import (
	"mini-project/config"
	"mini-project/models"
)

func GetUserById(id uint) (user models.User, err error) {
	if err := config.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return
}

func CountDonateUser(id uint) (result int) {
	var pet models.Pet
	if err := config.DB.Model(&pet).Where("user_id = ?", id).Count(&result).Error; err != nil {
		return 0
	}
	return
}

func CountAdoptUser(id uint) (result int) {
	var adopt models.Adoption
	if err := config.DB.Model(&adopt).Where("user_id = ?", id).Count(&result).Error; err != nil {
		return 0
	}
	return
}

func GetProfil(id uint) (user models.User, err error) {

	if err = config.DB.Preload("UserDetail").First(&user, id).Error; err != nil {
		return
	}

	return
}

func UpdateProfilDetail(req *models.UserDetail, id uint) error {
	var userDetail models.UserDetail
	if err := config.DB.Where("user_id = ?", id).First(&userDetail).Error; err != nil {
		if err := config.DB.Model(&userDetail).Where("user_id = ?", id).Save(&models.UserDetail{
			Alamat:    req.Alamat,
			Handphone: req.Handphone,
			UserID:    id,
		}).Error; err != nil {
			return err
		}
	} else {
		if err := config.DB.Model(&userDetail).Where("user_id = ?", id).Updates(models.UserDetail{
			Alamat:    req.Alamat,
			Handphone: req.Handphone,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func UpdateProfil(req *models.User, username string) error {

	if err := config.DB.Model(&req).Where("username = ?", username).Updates(models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id uint) error {
	var user models.User
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}
