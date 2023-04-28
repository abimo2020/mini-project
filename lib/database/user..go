package database

import (
	"mini-project/config"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo"
)

// func dashboard() (interface{}, error) {
// }

func GetProfil(c echo.Context) (interface{}, error) {
	var user models.User
	id := Authorization(c)

	if err := config.DB.Where("username = ?", id).Preload("UserDetail").First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
func UpdateProfilDetail(c echo.Context) (interface{}, error) {
	var userDetail models.UserDetail
	var user models.User

	id := Authorization(c)

	if err := config.DB.Where("username = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Where("user_id = ?", user.ID).First(&userDetail).Error; err != nil {
		c.Bind(&userDetail)
		if err := config.DB.Model(&userDetail).Where("user_id = ?", user.ID).Save(&models.UserDetail{
			Alamat:    userDetail.Alamat,
			Handphone: userDetail.Handphone,
			UserID:    user.ID,
		}).Error; err != nil {
			return nil, err
		}
	} else {
		c.Bind(&userDetail)
		if err := config.DB.Model(&userDetail).Where("user_id = ?", user.ID).Updates(models.UserDetail{
			Alamat:    userDetail.Alamat,
			Handphone: userDetail.Handphone,
		}).Error; err != nil {
			return nil, err
		}
	}
	return userDetail, nil
}

func UpdateProfil(c echo.Context) (interface{}, error) {
	var user models.User

	id := Authorization(c)

	password := c.FormValue("Password")
	if err := config.DB.Where("username = ? AND password = ?", id, password).First(&user).Error; err != nil {

		return nil, echo.NewHTTPError(http.StatusForbidden, "The password is wrong")
	}
	newName := c.FormValue("New Name")
	newEmail := c.FormValue("New Email")
	newPassword := c.FormValue("New Password")
	retypePassword := c.FormValue("Retype Password")
	if newPassword == retypePassword {
		if err := config.DB.Model(&user).Where("username = ?", id).Updates(models.User{
			Name:     newName,
			Email:    newEmail,
			Password: newPassword,
		}).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, echo.NewHTTPError(http.StatusForbidden, "The password is not match")
	}
	return user, nil
}

func DeleteUser(c echo.Context) (interface{}, error) {
	var user models.User

	id := Authorization(c)

	password := c.FormValue("Password")
	if err := config.DB.Where("username = ? AND password = ?", id, password).First(&user).Error; err != nil {

		return nil, echo.NewHTTPError(http.StatusForbidden, "The password is wrong")
	}

	if err := config.DB.Where("username = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// func DeletePet(c echo.Context) (interface{}, error) {
// 	id, _ := strconv.Atoi(c.Param("id"))
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
