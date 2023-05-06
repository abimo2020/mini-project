package usecase

import (
	"mini-project/lib/database"
	"mini-project/models"
	"mini-project/models/payload"
)

func Login(req *payload.LoginForm) (user models.User, err error) {
	u, e := database.Login(req)
	if e != nil {
		err = e
		return
	}
	user = u
	return
}

func Register(req *payload.Register) error {
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := database.Register(&user); err != nil {
		return err
	}
	return nil
}
