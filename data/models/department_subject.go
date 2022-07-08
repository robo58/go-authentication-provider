package models

import "github.com/jinzhu/gorm"

type DepartmentSubject struct {
	gorm.Model
	DepartmentId int `json:"department_id" gorm:"primaryKey"`
	SubjectId int `json:"subject_id" gorm:"primaryKey"`
	Students []DepartmentStudent `json:"students" gorm:"many2many:student_department_subjects;"`
}
