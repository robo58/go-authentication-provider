package schools

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/robo58/go-authentication-provider/data"
	"github.com/robo58/go-authentication-provider/data/models"
	"github.com/robo58/go-authentication-provider/helpers"
	"strings"
)

func GetSchools(c *gin.Context){
	db := data.GetDB()
	var schools []models.School
	scopes := strings.Split(sessions.Default(c).Get("scope").(string), " ")
	db = db.Preload("Headmaster")
	if helpers.Contains(scopes, "departments.read") {
		db = db.Preload("Departments.HeadroomTeacher")
	}
	if helpers.Contains(scopes, "subjects.read") {
		db = db.Preload("Departments.Subjects.Teacher")
	}
	if helpers.Contains(scopes, "students.read") {
		db = db.Preload("Departments.Students.Subjects.Subject.Teacher")
	}
	db.Find(&schools)
	c.JSON(200, gin.H{
		"schools": schools,
	})
}
