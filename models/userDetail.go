package models

import "github.com/jinzhu/gorm"

type UserDetail struct {
	gorm.Model
	Alamat    string `json:"alamat" form:"alamat" validate:"required"`
	Handphone string `json:"handphone" form:"handphone" validate:"required,min=11,max=13"`
	UserID    uint   `json:"user_id" form:"user_id" gorm:"unique"`
}
