package models

import "github.com/jinzhu/gorm"

type UserDetail struct {
	gorm.Model
	Alamat    string `json:"alamat" form:"alamat"`
	Handphone string `json:"handphone" form:"handphone"`
	UserID    uint   `json:"user_id" form:"user_id" gorm:"unique"`
}
