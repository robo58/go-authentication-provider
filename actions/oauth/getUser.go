package oauth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/robo58/go-authentication-provider/data"
	"github.com/robo58/go-authentication-provider/data/models"
	"net/http"
)

func GetUser(c *gin.Context) {
	var user models.User
	data.GetDB().Preload("Roles").First(&user, c.Query("user_id"))
	scopes := sessions.Default(c).Get("scope").(string)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"scopes": scopes,
		"token": c.GetHeader("authorization"),
	})
}

