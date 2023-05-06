package usecase

import (
	"mini-project/lib/database"
	"mini-project/models"
	"mini-project/models/payload"
	"net/http"

	"github.com/labstack/echo"
)

func DashboardAdmin() (response payload.DashboardAdmin) {
	totalUser := database.CountUsers()
	totalAdoption := database.CountAdoption()
	totalDonation := database.CountPets()
	totalAvailable := database.CountAvailablePets()
	totalAdopted := database.CountAdoptedPets()
	response = payload.DashboardAdmin{
		TotalUser:         totalUser,
		TotalAdoption:     totalAdoption,
		TotalDonation:     totalDonation,
		TotalPetAvailable: totalAvailable,
		TotalPetAdopted:   totalAdopted,
	}
	return response
}

func CreatePetCategory(req *payload.CreateCategory) error {
	category := models.PetCategory{
		Name: req.Name,
	}
	err := database.CreatePetCategory(&category)
	if err != nil {
		return err
	}
	return nil
}

func GetPetCategory() (response []string, err error) {
	category, err := database.GetPetCategory()
	if err != nil {
		err = echo.NewHTTPError(http.StatusBadRequest, "Don't have permission")
		return
	}
	for _, value := range category {
		response = append(response, value.Name)
	}
	return
}

func GetUsers() (response []payload.GetProfil, err error) {
	users, err := database.GetUsers()
	if err != nil {
		return
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
	return
}
