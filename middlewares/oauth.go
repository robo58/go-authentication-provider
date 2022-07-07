package middlewares

import (
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go"
	"github.com/robo58/go-authentication-provider/data"
	"log"
	"net/http"
	"strings"
)

func OauthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		log.Printf("authorization header: %s\n", authorizationHeader)
		if authorizationHeader!="" {
			token := strings.Split(authorizationHeader, " ")[1]
			admin := client.NewAPIClient(data.GetHydraConfig().Admin).AdminApi
			exec, _, err := admin.IntrospectOAuth2Token(c).Token(token).Execute()
			if err == nil && exec.GetActive() {
				c.AddParam("user_id",exec.GetSub())
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
}