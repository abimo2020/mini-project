package models

import "github.com/jinzhu/gorm"

type Adoption struct {
	gorm.Model
	UserID uint
	PetID  uint `gorm:"unique"`
}
