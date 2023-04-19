package models

import "github.com/jinzhu/gorm"

type PetCategory struct {
	gorm.Model
	Name string
}
