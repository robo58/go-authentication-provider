package models

import "github.com/jinzhu/gorm"

type RoleUser struct {
	gorm.Model
	UserId int `json:"user_id" gorm:"primaryKey"`
	RoleId int `json:"role_id" gorm:"primaryKey"`
}