package seeders

import (
	"github.com/jinzhu/gorm"
	"github.com/robo58/go-authentication-provider/data/models"
)

func CreateSchool(db *gorm.DB, name string, headmasterId int) error {
	return db.Create(&models.School{
		Name: name,
		HeadmasterId: headmasterId,
	}).Error
}
