package seeders

import (
	"github.com/jinzhu/gorm"
	"github.com/robo58/go-authentication-provider/data/models"
)

func CreateDepartment(db *gorm.DB, name string, schoolId int, headroomTeacherId int) error {
	return db.Create(&models.Department{
		Name: name,
		SchoolId: schoolId,
		HeadroomTeacherId: headroomTeacherId,
	}).Error
}

