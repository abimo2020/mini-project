package controller

// dasboard, see profil, update profil, delete user,
import (
	"net/http"

	"mini-project/lib/database"

	"github.com/labstack/echo"
)

func GetProfilController(c echo.Context) error {
	user, err := database.GetProfil(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all pets",
		"user":    user,
	})
}

func UpdateProfilDetailController(c echo.Context) error {

	user, err := database.UpdateProfilDetail(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update profil user",
		"user":   user,
	})
}

func UpdateProfilController(c echo.Context) error {
	user, err := database.UpdateProfil(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update user",
		"user":   user,
	})
}

func DeleteUserController(c echo.Context) error {
	user, err := database.DeleteUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete user",
		"user":   user,
	})
}

// // get all pets
// func GetPetsController(c echo.Context) error {
// 	pets, err := database.GetPets()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get all pets",
// 		"pets":    pets,
// 	})
// }

// // get pet by id
// func GetPetController(c echo.Context) error {
// 	pet, err := database.GetPet(c)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get all pets",
// 		"pets":    pet,
// 	})
// }

// // delete pet by id
// func DeletePetController(c echo.Context) error {
// 	pet, err := database.DeletePet(c)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success delete pet",
// 		"pet":     pet,
// 	})
// }
// func CreatePetController(c echo.Context) error {
// 	pet, err := database.CreatePet(c)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success create new pet",
// 		"pet":     pet,
// 	})
// }

// // update pet by id
// // func UpdatePetController(c echo.Context) error {
// // 	pet, err := database.UpdatePet(c)
// // 	if err != nil {
// // 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// // 	}
// // 	return c.JSON(http.StatusOK, map[string]interface{}{
// // 		"message": "success update pet",
// // 		"pet":     pet,
// // 	})

// // }