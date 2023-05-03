package controller

import (
	"net/http"

	"mini-project/lib/database"

	"github.com/labstack/echo"
)

// get all pets
func GetPetsController(c echo.Context) error {
	pets, err := database.GetPets()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all pets",
		"pets":    pets,
	})
}

// get pet by id
func GetPetController(c echo.Context) error {
	pet, err := database.GetPetDetail(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get the pet detail",
		"pets":    pet,
	})
}

// delete pet by id
func DeletePetController(c echo.Context) error {
	err := database.DeletePet(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete pet",
	})
}

func DonateController(c echo.Context) error {
	err := database.DonatePet(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success donate pet",
	})
}

func AdoptController(c echo.Context) error {
	owner, err := database.AdoptPet(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success adopt pet",
		"owner":   owner,
	})
}

func GetDonateListController(c echo.Context) error {
	donate, err := database.GetDonateList(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get donate list",
		"donate":  donate,
	})
}

func GetAdoptListController(c echo.Context) error {
	adopt, err := database.GetAdoptList(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get adopt list",
		"adopt":   adopt,
	})
}

func UpdatePetStatusController(c echo.Context) error {
	err := database.UpdatePetStatus(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update pet status",
	})
}

func UpdatePetController(c echo.Context) error {
	err := database.UpdatePet(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update pet",
	})
}
