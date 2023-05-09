package usecase

import (
	"mini-project/lib/database"
	"mini-project/models/payload"
	"net/http"

	"github.com/labstack/echo"
)

func GetAvailablePets() (response []payload.GetPet, err error) {
	pets, err := database.GetAvailablePets()
	if err != nil {
		return []payload.GetPet{}, err
	}

	for _, value := range pets {
		response = append(response, payload.GetPet{
			ID:        value.ID,
			Deskripsi: value.Deskripsi,
			Status:    value.Status,
			Category:  value.PetCategory.Name,
			OwnerID:   value.UserID,
			Owner:     value.User.Name,
			Handphone: value.User.UserDetail.Handphone,
			Alamat:    value.User.UserDetail.Alamat,
		})
	}
	return response, nil
}

func GetPet(id uint) (response payload.GetPet, err error) {
	pet, err := database.GetPet(id)
	if err != nil {
		return payload.GetPet{}, err
	}

	response = payload.GetPet{
		ID:        pet.ID,
		Deskripsi: pet.Deskripsi,
		Status:    pet.Status,
		Category:  pet.PetCategory.Name,
		OwnerID:   pet.UserID,
		Owner:     pet.User.Name,
		Handphone: pet.User.UserDetail.Handphone,
		Alamat:    pet.User.UserDetail.Alamat,
	}
	return response, nil
}

func GetPets() (response []payload.GetPet, err error) {
	pets, err := database.GetPets()
	if err != nil {
		return
	}
	for _, value := range pets {
		response = append(response, payload.GetPet{
			ID:        value.ID,
			Deskripsi: value.Deskripsi,
			Status:    value.Status,
			Category:  value.PetCategory.Name,
			OwnerID:   value.UserID,
			Owner:     value.User.Name,
			Handphone: value.User.UserDetail.Handphone,
			Alamat:    value.User.UserDetail.Alamat,
		})
	}
	return
}

func GetDonates(id uint) (response []payload.GetDonateList, err error) {
	user, err := database.GetDonateList(id)
	if err != nil {
		return
	}
	for _, value := range user.Pet {
		response = append(response, payload.GetDonateList{
			ID:        value.ID,
			Deskripsi: value.Deskripsi,
			Status:    value.Status,
			Category:  value.PetCategory.Name,
		})
	}
	return
}

func GetAdoptions(id uint) (response []payload.GetPet, err error) {
	user, err := database.GetAdoptList(id)
	if err != nil {
		return
	}
	for _, value := range user.Adoption {
		response = append(response, payload.GetPet{
			Deskripsi: value.Pet.Deskripsi,
			Category:  value.Pet.PetCategory.Name,
			OwnerID:   value.UserID,
			Owner:     value.Pet.User.Name,
			Handphone: value.Pet.User.UserDetail.Handphone,
			Alamat:    value.Pet.User.UserDetail.Alamat,
		})
	}
	return
}

func UpdatePet(req *payload.UpdatePet, id uint, petId uint) error {
	pet, err := database.GetPet(petId)
	if err != nil {
		return nil
	}
	if pet.UserID != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Don't have permission")
	}
	if pet.Status == "adopted" {
		return echo.NewHTTPError(http.StatusBadRequest, "Can't update adopted pet")
	}
	if err := database.UpdatePet(req, petId); err != nil {
		return err
	}
	return nil
}
func UpdateAvailableStatus(id uint, petId uint) error {
	pet, err := database.GetPet(petId)
	if err != nil {
		return err
	}
	if pet.UserID != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Don't have permission")
	}
	if err := database.UpdateAvailableStatus(pet.ID); err != nil {
		return err
	}
	return nil
}

func DonatePet(req *payload.CreatePet, id uint) error {
	profil, err := database.GetProfil(id)
	if err != nil {
		return err
	}
	if profil.UserDetail.UserID != id {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Fill the address and handphone before adopt or donate")
	}
	if err := database.DonatePet(req, id); err != nil {
		return err
	}
	return nil
}

func AdoptPet(id uint, petId uint) (response payload.GetPet, err error) {
	profil, err := database.GetProfil(id)
	if err != nil {
		return
	}
	if profil.UserDetail.UserID != id {
		return response, echo.NewHTTPError(http.StatusNotAcceptable, "Fill the address and handphone before adopt or donate")
	}
	pet, err := database.GetPet(petId)
	if err != nil {
		return
	}
	if pet.UserID == id {
		return response, echo.NewHTTPError(http.StatusNotAcceptable, "Owner can't adopt the pet")
	}
	if err := database.UpdateAdoptedStatus(petId); err != nil {
		return response, err
	}
	if err := database.AdoptPet(id, petId); err != nil {
		return response, err
	}
	response = payload.GetPet{
		Deskripsi: pet.Deskripsi,
		Category:  pet.PetCategory.Name,
		OwnerID:   pet.UserID,
		Owner:     pet.User.Name,
		Handphone: pet.User.UserDetail.Handphone,
		Alamat:    pet.User.UserDetail.Alamat,
	}
	return
}

func DeletePet(id uint, petId uint) error {
	pet, err := database.GetPet(petId)
	if err != nil {
		return err
	}
	if pet.UserID != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Don't have permission")
	}
	if pet.Status == "adopted" {
		return echo.NewHTTPError(http.StatusBadRequest, "Can't delete adopted pet")
	}
	if err := database.DeletePet(petId); err != nil {
		return err
	}
	return nil
}
