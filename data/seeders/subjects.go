package seeders

import (
	"github.com/jinzhu/gorm"
	"github.com/robo58/go-authentication-provider/data/models"
)

func CreateSubject(db *gorm.DB, name string, teacherId int) error {
	return db.Create(&models.Subject{
		Name: name,
		TeacherId: teacherId,
	}).Error
}
