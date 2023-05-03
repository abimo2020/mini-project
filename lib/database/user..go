package database

import (
	"mini-project/config"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo"
)

func DashboardUser(c echo.Context) (interface{}, interface{}, error) {
	var pet models.Pet
	var adopt models.Adoption
	var user_donate int
	var user_adopt int

	_, id := Authorization(c)

	if err := config.DB.Model(&pet).Where("user_id = ?", id).Count(&user_donate).Error; err != nil {
		user_donate = 0
	}
	if err := config.DB.Model(&adopt).Where("user_id = ?", id).Count(&user_adopt).Error; err != nil {
		user_adopt = 0
	}
	return user_donate, user_adopt, nil
}

func GetProfil(c echo.Context) (interface{}, error) {
	var user models.User
	username, _ := Authorization(c)

	if err := config.DB.Where("username = ?", username).Preload("UserDetail").First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateProfilDetail(c echo.Context) error {
	var userDetail models.UserDetail
	var user models.User

	username, _ := Authorization(c)

	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}
	if err := config.DB.Where("user_id = ?", user.ID).First(&userDetail).Error; err != nil {
		c.Bind(&userDetail)
		if err := config.DB.Model(&userDetail).Where("user_id = ?", user.ID).Save(&models.UserDetail{
			Alamat:    userDetail.Alamat,
			Handphone: userDetail.Handphone,
			UserID:    user.ID,
		}).Error; err != nil {
			return err
		}
	} else {
		c.Bind(&userDetail)
		if err := config.DB.Model(&userDetail).Where("user_id = ?", user.ID).Updates(models.UserDetail{
			Alamat:    userDetail.Alamat,
			Handphone: userDetail.Handphone,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func UpdateProfil(c echo.Context) error {
	var user models.User

	username, _ := Authorization(c)

	password := c.FormValue("Password")
	if err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, "The password is wrong")
	}
	newName := c.FormValue("New Name")
	newEmail := c.FormValue("New Email")
	newPassword := c.FormValue("New Password")
	retypePassword := c.FormValue("Retype Password")
	if newPassword == retypePassword {
		if err := config.DB.Model(&user).Where("username = ?", username).Updates(models.User{
			Name:     newName,
			Email:    newEmail,
			Password: newPassword,
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

// func DeletePet(c echo.Context) (interface{}, error) {
// 	username, _ := strconv.Atoi(c.Param("id"))
// 	var pet models.Pet
// 	if err := config.DB.First(&pet, id).Error; err != nil {
// 		return nil, err
// 	}
// 	if err := config.DB.Delete(&pet, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return pet, nil
// }

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
