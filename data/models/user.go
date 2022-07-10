package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Roles 		[]*Role `json:"roles,omitempty" gorm:"many2many:role_users;"`
}

type ResourceUser struct {
	ID 			uint `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Roles 		[]*Role `json:"roles,omitempty" gorm:"many2many:role_users;"`
}