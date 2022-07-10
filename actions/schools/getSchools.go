package schools

import (
	"github.com/gin-gonic/gin"
	"github.com/robo58/go-authentication-provider/data"
	"github.com/robo58/go-authentication-provider/data/models"
)

func GetSchools(c *gin.Context){
	db := data.GetDB()
	var schools []models.School
	db.Preload("Headmaster.Roles").Preload("Departments.HeadroomTeacher").Preload("Departments.Subjects.Teacher").Preload("Departments.Students.Subjects").Find(&schools)
	c.JSON(200, gin.H{
		"schools": schools,
	})
}
