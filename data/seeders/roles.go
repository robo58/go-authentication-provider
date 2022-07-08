package seeders

import (
	"github.com/jinzhu/gorm"
	"github.com/robo58/go-authentication-provider/data/models"
)

func CreateRole(db *gorm.DB, name string) error {
	return db.Create(&models.Role{
		Name: name,
	}).Error
}
