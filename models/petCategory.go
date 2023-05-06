package models

import "github.com/jinzhu/gorm"

type PetCategory struct {
	gorm.Model
	Name string `json:"name" form:"name" gorm:"unique" validate:"required"`
}
