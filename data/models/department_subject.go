package models

import "github.com/jinzhu/gorm"

type DepartmentSubject struct {
	gorm.Model
	DepartmentId int `json:"department_id" gorm:"primaryKey"`
	SubjectId int `json:"subject_id" gorm:"primaryKey"`
	Subject *Subject `json:"subject,omitempty" gorm:"foreignKey:SubjectId"`
	Department *Department `json:"department,omitempty" gorm:"foreignKey:DepartmentId"`
	Students []*DepartmentStudent `json:"students,omitempty" gorm:"many2many:student_department_subjects;"`
	// pivot fields
	Grade int `json:"grade" gorm:"-"`

}
