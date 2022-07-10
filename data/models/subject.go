package models

import "github.com/jinzhu/gorm"

type Subject struct {
	gorm.Model
	Name        string       `json:"name"`
	TeacherId   int          `json:"teacher_id"`
	Teacher     *User         `json:"teacher,omitempty" gorm:"foreignKey:TeacherId;"`
	Departments []*Department `json:"departments,omitempty" gorm:"many2many:department_subjects;"`
}
