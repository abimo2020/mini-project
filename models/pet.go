package models

import "github.com/jinzhu/gorm"

type Pet struct {
	gorm.Model
	Category      PetCategory
	Deskripsi     string `json:"deskripsi" form:"deskripsi" gorm:"not null"`
	Status        bool   `json:"status" form:"status" gorm:"default:false"`
	UserID        uint   `json:"user_id" form:"user_id"`
	PetCategoryID uint   `json:"pet_category_id" form:"pet_category_id"`
}
