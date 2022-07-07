package seeders

import (
	"github.com/jinzhu/gorm"
)

type Seeder struct {
	Name string
	Run func(*gorm.DB) error
}

func All() []Seeder {
	return []Seeder{
		Seeder{
			Name: "CreateUserJane",
			Run: func(db *gorm.DB) error {
				err := CreateUser(db, "Jane", "jane.doe@mail.com", "password")
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}