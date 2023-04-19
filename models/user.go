package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Username     string `gorm:"unique"`
	Email        string `json:"email" form:"email" gorm:"unique`
	Password     string `json:"password" form:"password"`
	UserDetailID uint   `gorm:"unique`
	Token        string `gorm:"-"`
}

// func (u *User) AfterDelete(tx *gorm.DB) (err error) {
// 	tx.Where("user_id = ?", u.ID).Delete(&Blog{})
// 	return
// }
