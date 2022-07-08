package models

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	Name              string             `json:"name"`
	SchoolId int `json:"school_id"`
	HeadroomTeacherId int                `json:"headroom_teacher_id"`
	School            School             `gorm:"foreignKey:SchoolId;"`
	HeadroomTeacher   User               `gorm:"foreignKey:HeadroomTeacherId;"`
	Subjects          []Subject `json:"subjects" gorm:"many2many:department_subjects;"`
	Students          []User `json:"students" gorm:"many2many:department_students;"`
}