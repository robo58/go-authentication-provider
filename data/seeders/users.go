package seeders

import (
	"github.com/jinzhu/gorm"
	"github.com/robo58/go-authentication-provider/data/models"
	"github.com/robo58/go-authentication-provider/helpers/crypto"
)

func CreateUser(db *gorm.DB, name string, email string, password string) error {
	return db.Create(&models.User{
		Name: name,
		Email: email,
		Password:      crypto.HashAndSalt([]byte(password)),
	}).Error
}