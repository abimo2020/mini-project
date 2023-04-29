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
	err := database.AdoptPet(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success adopt pet",
	})
}
