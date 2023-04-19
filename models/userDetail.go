package models

import "github.com/jinzhu/gorm"

type UserDetail struct {
	gorm.Model
	Alamat    string
	Handphone string
}
