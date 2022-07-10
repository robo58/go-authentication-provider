package models

import "github.com/jinzhu/gorm"

type School struct {
	gorm.Model
	Name string `json:"name"`
	HeadmasterId int `json:"headmaster_id"`
	Headmaster *User `json:"headmaster,omitempty" gorm:"foreignKey:HeadmasterId;"`
	Departments []*Department `json:"departments,omitempty" gorm:"foreignKey:SchoolId;"`
}