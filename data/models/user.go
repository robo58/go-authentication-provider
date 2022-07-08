package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Roles []Role       `json:"roles" gorm:"many2many:role_users;"`
}
