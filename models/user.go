package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name       string `json:"name" form:"name" gorm:"not null" validate:"required"`
	Username   string `json:"username" form:"username" gorm:"unique;not null" validate:"required,min=4,max=12"`
	Role       string `json:"role" form:"role" gorm:"type:enum('user','admin');default:'user'"`
	Email      string `json:"email" form:"email" gorm:"unique;not null" validate:"required,email"`
	Password   string `json:"password" form:"password" validate:"required,min=6,max=12"`
	UserDetail UserDetail
	Adoption   []Adoption
	Pet        []Pet
}

type Token struct {
	Token string `gorm:"-"`
}

// func (u *User) AfterDelete(tx *gorm.DB) (err error) {
// 	tx.Where("user_id = ?", u.ID).Delete(&Blog{})
// 	return
// }
