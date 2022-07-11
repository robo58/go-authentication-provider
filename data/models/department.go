package models

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	Name              string             `json:"name"`
	SchoolId int `json:"school_id"`
	HeadroomTeacherId int                `json:"headroom_teacher_id"`
	School            *School             `json:"school,omitempty" gorm:"foreignKey:SchoolId;"`
	HeadroomTeacher   *User               `json:"headroom_teacher,omitempty;" gorm:"foreignKey:HeadroomTeacherId;"`
	Subjects          []*Subject `json:"subjects,omitempty" gorm:"many2many:department_subjects;"`
	Students          []*DepartmentStudent `json:"students,omitempty" gorm:"foreignKey:DepartmentId;"`
}