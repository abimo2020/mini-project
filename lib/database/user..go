package database

import (
	"mini-project/config"
	"mini-project/models"
	"mini-project/models/payload"
	"net/http"

	"github.com/labstack/echo"
)

func DashboardUser(c echo.Context) (interface{}, error) {
	var pet models.Pet
	var adopt models.Adoption
	var dashboard payload.DashboardUser

	_, id := Authorization(c)

	if err := config.DB.Model(&pet).Where("user_id = ?", id).Count(&dashboard.Donation).Error; err != nil {
		dashboard.Donation = 0
	}
	if err := config.DB.Model(&adopt).Where("user_id = ?", id).Count(&dashboard.Adoption).Error; err != nil {
		dashboard.Adoption = 0
	}
	return dashboard, nil
}

func GetProfil(c echo.Context) (interface{}, error) {
	var user models.User

	username, _ := Authorization(c)

	if err := config.DB.Where("username = ?", username).Preload("UserDetail").First(&user).Error; err != nil {
		return nil, err
	}

	profil := payload.GetProfil{
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Alamat:    user.UserDetail.Alamat,
		Handphone: user.UserDetail.Handphone,
	}

	return profil, nil
}

func UpdateProfilDetail(c echo.Context) error {
	var userDetail models.UserDetail
	var user models.User
	var update payload.UpdateProfilDetail

	username, _ := Authorization(c)

	c.Bind(&update)

	if err := c.Validate(&update); err != nil {
		return err
	}

	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}
	if err := config.DB.Where("user_id = ?", user.ID).First(&userDetail).Error; err != nil {
		if err := config.DB.Model(&userDetail).Where("user_id = ?", user.ID).Save(&models.UserDetail{
			Alamat:    update.Alamat,
			Handphone: update.Handphone,
			UserID:    user.ID,
		}).Error; err != nil {
			return err
		}
	} else {
		if err := config.DB.Model(&userDetail).Where("user_id = ?", user.ID).Updates(models.UserDetail{
			Alamat:    update.Alamat,
			Handphone: update.Handphone,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func UpdateProfil(c echo.Context) error {
	var user models.User
	var updateProfil payload.UpdateProfil
	username, _ := Authorization(c)

	c.Bind(&updateProfil)

	if err := c.Validate(updateProfil); err != nil {
		return err
	}

	if err := config.DB.Where("username = ? AND password = ?", username, updateProfil.Password).First(&user).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, "The password is wrong")
	}

	if updateProfil.NewPassword == updateProfil.RetypePassword {
		if err := config.DB.Model(&user).Where("username = ?", username).Updates(models.User{
			Name:     updateProfil.Name,
			Email:    updateProfil.Email,
			Password: updateProfil.NewPassword,
		}).Error; err != nil {
			return err
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "The password is not match")
	}
	return nil
}

func DeleteUser(c echo.Context) error {
	var user models.User

	username, _ := Authorization(c)

	password := c.FormValue("Password")
	if err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {

		return echo.NewHTTPError(http.StatusForbidden, "The password is wrong")
	}

	if err := config.DB.Where("username = ?", username).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
