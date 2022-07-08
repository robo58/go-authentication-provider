package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/robo58/go-authentication-provider/config"
	"github.com/robo58/go-authentication-provider/data/models"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() {
	var db = DB

	configuration := config.GetConfig()

	driver := configuration.Database.Driver
	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	if driver == "sqlite" { // SQLITE
		db, err = gorm.Open("sqlite3", "./"+database+".db")
		if err != nil {
			fmt.Println("db err: ", err)
		}
	} else if driver == "postgres" { // POSTGRES
		db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
		if err != nil {
			fmt.Println("db err: ", err)
		}
	} else if driver == "mysql" { // MYSQL
		db, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println("db err: ", err)
		}
	}

	// Change this to true if you want to see SQL queries
	db.LogMode(false)
	db.DB().SetMaxIdleConns(configuration.Database.MaxIdleConns)
	db.DB().SetMaxOpenConns(configuration.Database.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)
	DB = db
	migration()
}

// Auto migrate project models
func migration() {
	DB.AutoMigrate(&models.StudentDepartmentSubject{})
	DB.AutoMigrate(&models.RoleUser{})
	DB.AutoMigrate(&models.DepartmentSubject{})
	DB.AutoMigrate(&models.DepartmentStudent{})
	DB.AutoMigrate(&models.School{})
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Subject{})
	DB.AutoMigrate(&models.Department{})
}

func GetDB() *gorm.DB {
	return DB
}
