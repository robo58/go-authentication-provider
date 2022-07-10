package middlewares

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go"
	"github.com/robo58/go-authentication-provider/data"
	"log"
	"net/http"
	"strings"
)



func AccessTokenRequired() gin.HandlerFunc {
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

func GetTokenFromSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("GetTokenFromSession")
		session := sessions.Default(c)
		token := session.Get("access_token")
		fmt.Println("token: ", token)
		if token != nil {
			c.Request.Header.Set("authorization", "Bearer "+ token.(string))
		}
		c.Next()
	}
}

func ScopesRequired(scopes []string) gin.HandlerFunc  {
	return func(c *gin.Context) {
		fmt.Println("ScopesRequired")
		fmt.Println("scopes: ", scopes)
		session := sessions.Default(c)
		sessionScope := session.Get("scope")
		fmt.Println("session scopes: ", sessionScope)
		if len(scopes)>0 {
			if  sessionScope != nil {
				arrayScopes := strings.Split(sessionScope.(string), " ")
				var missingScopes []string
				for _, scope := range scopes {
					if !contains(arrayScopes, scope) {
						missingScopes = append(missingScopes, scope)
					}
				}
				if len(missingScopes) > 0 {
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden", "missingScopes": missingScopes})
					return
				}
			} else {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden", "missingScopes": scopes})
				return
			}
		}
		c.Next()
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}