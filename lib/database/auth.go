package database

import (
	"mini-project/config"
	"mini-project/lib/cookie"
	"mini-project/middlewares"
	"mini-project/models"
	"mini-project/models/payload"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	var err error
	user := models.User{}
	login := payload.LoginForm{}
	c.Bind(&login)

	if err := c.Validate(login); err != nil {
		return err
	}

	_, e := c.Cookie("JWTCookie")

	if e == nil {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "Already logged in")
	}

	if err = config.DB.Where("username = ?", login.ID).Where("password = ? ", login.Password).First(&user).Error; err != nil {
		if err = config.DB.Where("email = ?", login.ID).Where("password = ? ", login.Password).First(&user).Error; err != nil {
			return err
		}
	}
	token, err := middlewares.CreateToken(user.Username, user.Role, user.ID)

	cookie.CreateJWTCookies(c, token)

	if err != nil {
		return err
	}
	return nil
}

func Logout(c echo.Context) error {
	cookie, err := c.Cookie("JWTCookie")
	if err != nil {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "Not logged in yet")
	}
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	c.SetCookie(cookie)
	return nil
}

func Register(c echo.Context) error {

	register := payload.Register{}
	_, e := c.Cookie("JWTCookie")

	if e == nil {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "Already logged in")
	}
	c.Bind(&register)

	if err := c.Validate(register); err != nil {
		return err
	}

	if register.Password != register.RetypePassword {
		return echo.NewHTTPError(http.StatusBadRequest, "Password not match")
	}

	user := models.User{
		Name:     register.Name,
		Username: register.Username,
		Email:    register.Email,
		Password: register.Password,
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func Authorization(c echo.Context) (string, uint) {
	cookie, _ := c.Cookie("JWTCookie")
	token, _ := jwt.Parse(cookie.Value, nil)
	claims, _ := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	id := uint(claims["user_id"].(float64))
	return username, id
}
