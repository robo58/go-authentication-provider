package models

import "github.com/jinzhu/gorm"

type DepartmentStudent struct {
	gorm.Model
	DepartmentId int        `json:"department_id" gorm:"primaryKey"`
	UserId    int        `json:"user_id" gorm:"primaryKey"`
	Student      *User       `json:"student,omitempty" gorm:"foreignKey:UserId;"`
	Department   *Department `json:"department,omitempty" gorm:"foreignKey:DepartmentId;"`
	Subjects []*DepartmentSubject `json:"subjects,omitempty" gorm:"many2many:student_department_subjects;"`
}
