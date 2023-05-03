package models

import "github.com/jinzhu/gorm"

type Pet struct {
	gorm.Model
	UserID        uint   `json:"user_id" form:"user_id"`
	Deskripsi     string `json:"deskripsi" form:"deskripsi" gorm:"not null"`
	Status        string `json:"status" form:"status" gorm:"type:enum('adopted','available');default:'available'"`
	PetCategoryID uint   `json:"pet_category_id" form:"pet_category_id"`
	PetCategory   PetCategory
}
