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

	if err := config.DB.Preload("PetCategory").Where("status = ?", "available").Find(&pets).Error; err != nil {
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

func GetDonateList(c echo.Context) (interface{}, error) {
	var user models.User
	_, id := Authorization(c)

	if err := config.DB.Preload("Pet").Preload("Pet.PetCategory").First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}
func GetAdoptList(c echo.Context) (interface{}, error) {
	var user models.User
	_, id := Authorization(c)

	if err := config.DB.Preload("Adoption").First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdatePet(c echo.Context) error {
	var pet models.Pet
	_, id := Authorization(c)
	petId := c.Param("id")

	if err := config.DB.First(&pet, petId).Error; err != nil {
		return err
	}

	if pet.UserID != id {
		return echo.NewHTTPError(http.StatusBadRequest, "Don't have permission")
	}

	c.Bind(&pet)

	if err := config.DB.Model(&pet).Where("user_id = ? AND id = ?", id, petId).Update(models.Pet{
		Deskripsi: pet.Deskripsi,
	}).Error; err != nil {
		return err
	}

	return nil
}

func UpdatePetStatus(c echo.Context) error {
	var pet models.Pet
	_, id := Authorization(c)
	petId := c.Param("id")

	if err := config.DB.First(&pet, petId).Error; err != nil {
		return err
	}

	if pet.UserID != id {
		return echo.NewHTTPError(http.StatusBadRequest, "Don't have permission")
	}

	if pet.Status == "available" {
		if err := config.DB.Model(&pet).Where("user_id = ? AND id = ?", id, petId).Update(models.Pet{
			Status: "adopted",
		}).Error; err != nil {
			return err
		}
	} else {
		if err := config.DB.Model(&pet).Where("user_id = ? AND id = ?", id, petId).Update(models.Pet{
			Status: "available",
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func DonatePet(c echo.Context) error {
	var pet models.Pet
	var profil models.UserDetail

	_, id := Authorization(c)
	c.Bind(&pet)

	if err := config.DB.Where("user_id = ?", id).First(&profil).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Fill the address and handphone before adopt or donate")
	}
	if err := config.DB.Model(&pet).Create(&models.Pet{
		Deskripsi:     pet.Deskripsi,
		PetCategoryID: pet.PetCategoryID,
		UserID:        id,
	}).Error; err != nil {
		return err
	}
	return nil
}

func AdoptPet(c echo.Context) (interface{}, error) {
	var pet models.Pet
	var profil models.UserDetail
	var adoption models.Adoption
	var user models.User

	_, user_id := Authorization(c)
	id, _ := strconv.Atoi(c.Param("id"))
	pet_id := uint(id)

	if err := config.DB.Where("user_id = ?", user_id).First(&profil).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Fill the address and handphone before adopt or donate")
	}

	if err := config.DB.Find(&pet, pet_id).Error; err != nil {
		return nil, err
	}

	if pet.UserID == user_id {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Owner can't adopt the pet")
	}

	if err := config.DB.Model(&pet).Where("id = ? AND status = ?", pet_id, "available").Update(models.Pet{
		Status: "adopted",
	}).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Model(&adoption).Save(&models.Adoption{
		UserID: user_id,
		PetID:  pet_id,
	}).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Preload("UserDetail").First(&user, pet.UserID).Error; err != nil {
		return nil, err
	}
	return user, nil
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
