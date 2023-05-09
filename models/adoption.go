package models

import "github.com/jinzhu/gorm"

type Adoption struct {
	gorm.Model
	UserID uint `json:"user_id" form:"user_id"`
	User   User
	PetID  uint `json:"pet_id" form:"pet_id" gorm:"unique"`
	Pet    Pet
}
