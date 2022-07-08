package models

import "github.com/jinzhu/gorm"

type StudentDepartmentSubject struct {
	gorm.Model
	DepartmentStudentId int `json:"department_student_id" gorm:"primaryKey"`
	DepartmentSubjectId int `json:"department_subject_id" gorm:"primaryKey"`
	Grade int `json:"grade"`
	DepartmentStudent DepartmentStudent `json:"department_student" gorm:"foreignKey:DepartmentStudentId;"`
	DepartmentSubject DepartmentSubject `json:"department_subject" gorm:"foreignKey:DepartmentSubjectId;"`
}
