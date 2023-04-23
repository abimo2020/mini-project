package middlewares

import (
	"mini-project/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["role"] = role
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_KEY))
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["role"].(string)
		if isAdmin != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
