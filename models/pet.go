package models

import "github.com/jinzhu/gorm"

type Pet struct {
	gorm.Model
	Category      PetCategory
	Deskripsi     string
	Status        bool
	UserID        uint
	PetCategoryID uint
}
