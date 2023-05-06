package usecase

import (
	"mini-project/lib/database"
	"mini-project/models"
	"mini-project/models/payload"
	"net/http"

	"github.com/labstack/echo"
)

func GetDashboardUser(id uint) payload.DashboardUser {
	donate := database.CountDonateUser(id)
	adopt := database.CountAdoptUser(id)
	response := payload.DashboardUser{
		Adoption: adopt,
		Donation: donate,
	}
	return response
}

func GetProfil(id uint) (payload.GetProfil, error) {
	user, err := database.GetProfil(id)
	if err != nil {
		return payload.GetProfil{}, err
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

func UpdateProfilDetail(req *payload.UpdateProfilDetail, id uint) error {
	profil := models.UserDetail{
		Alamat:    req.Alamat,
		Handphone: req.Handphone,
	}
	if err := database.UpdateProfilDetail(&profil, id); err != nil {
		return err
	}
	return nil
}

func UpdateProfil(id uint, req *payload.UpdateProfil) error {
	profil := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.NewPassword,
	}
	user, err := database.GetUserById(id)
	if err != nil {
		return err
	}
	if user.Password != req.Password {
		return echo.NewHTTPError(http.StatusBadRequest, "The password is wrong")
	}
	if req.NewPassword != req.RetypePassword {
		return echo.NewHTTPError(http.StatusBadRequest, "The password is not match")
	}

	if err := database.UpdateProfil(&profil, user.Username); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint, password string) error {
	user, err := database.GetUserById(id)
	if err != nil {
		return err
	}
	if user.Password != password {
		return echo.NewHTTPError(http.StatusBadRequest, "The password is wrong")
	}
	if err := database.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
