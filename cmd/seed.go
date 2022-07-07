package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/robo58/go-authentication-provider/config"
	db"github.com/robo58/go-authentication-provider/data"
	"github.com/robo58/go-authentication-provider/data/seeders"
	"log"
)

func main() {
	config.Setup("./config.yml")
	db.SetupDB()
	dbConn := db.GetDB()
	defer func(dbConn *gorm.DB) {
		err := dbConn.Close()
		if err != nil {
			return
		}
	}(dbConn)

	for _, seed := range seeders.All() {
		if err := seed.Run(dbConn); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}