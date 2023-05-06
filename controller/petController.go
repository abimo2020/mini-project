package controller

import (
	"mini-project/models/payload"
	"mini-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all pets
func GetPetsController(c echo.Context) error {
	pets, err := usecase.GetAvailablePets()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all pets",
		"data":    pets,
	})
}

// get pet by id
func GetPetController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	pet, err := usecase.GetPet(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get the pet detail",
		"data":    pet,
	})
}

// delete pet by id
func DeletePetController(c echo.Context) error {
	petId, _ := strconv.Atoi(c.Param("id"))
	_, id := Authorization(c)

	err := usecase.DeletePet(id, uint(petId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete pet",
	})
}

func DonateController(c echo.Context) error {
	var req payload.CreatePet
	_, id := Authorization(c)

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.DonatePet(&req, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success donate pet",
	})
}

func AdoptController(c echo.Context) error {
	_, id := Authorization(c)
	petId, _ := strconv.Atoi(c.Param("id"))
	response, err := usecase.AdoptPet(id, uint(petId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success adopt pet",
		"data":    response,
	})
}

func GetDonateListController(c echo.Context) error {
	_, id := Authorization(c)
	donate, err := usecase.GetDonates(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get donate list",
		"data":    donate,
	})
}

func GetAdoptListController(c echo.Context) error {
	_, id := Authorization(c)
	adopt, err := usecase.GetAdoptions(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get adopt list",
		"data":    adopt,
	})
}

func UpdatePetStatusController(c echo.Context) error {
	_, id := Authorization(c)
	petId, _ := strconv.Atoi(c.Param("id"))
	err := usecase.UpdatePetStatus(id, uint(petId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update pet status",
	})
}

func UpdatePetController(c echo.Context) error {
	var req payload.UpdatePet
	_, id := Authorization(c)
	petId, _ := strconv.Atoi(c.Param("id"))

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return err
	}

	if err := usecase.UpdatePet(&req, id, uint(petId)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update pet",
	})
}
