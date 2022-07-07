package users

import (
	"github.com/robo58/go-authentication-provider/data/models"
)

type User struct {
	models.Base
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
