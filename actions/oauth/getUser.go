package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/robo58/go-authentication-provider/data"
	"github.com/robo58/go-authentication-provider/data/models"
	"net/http"
)

func GetUser(c *gin.Context) {
	var user models.User
	data.GetDB().First(&user, c.Query("user_id"))

	c.JSON(http.StatusOK, user)
}

