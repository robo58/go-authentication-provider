package router

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robo58/go-authentication-provider/actions/oauth"
	"github.com/robo58/go-authentication-provider/middlewares"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"os"
)

// Endpoint is OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "http://localhost:4444/oauth2/auth",
	TokenURL: "http://localhost:4444/oauth2/token",
}

// OAuthConf Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var OAuthConf = &oauth2.Config{
	RedirectURL:  os.Getenv("REDIRECT_URL"),
	ClientID:     os.Getenv("CLIENT_ID"),     // TODO from hydra
	ClientSecret: os.Getenv("CLIENT_SECRET"), // TODO from hydra

	// https://github.com/coreos/go-oidc/blob/v3/oidc/oidc.go#L23-L36
	// offline scope for requesting Refresh Token
	// openid for Open ID Connect
	Scopes:   []string{"users.write", "users.read", "users.edit", "users.delete", "offline"},
	Endpoint: Endpoint,
}

var stateStore = map[string]bool{}


func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	app.LoadHTMLGlob("views/*")

	//// Routes
	app.GET("/oauth/login", oauth.GetLogin)
	app.POST("/oauth/login", oauth.PostLogin)
	app.GET("/oauth/consent", oauth.GetConsent)
	app.POST("/oauth/consent", oauth.PostConsent)
	app.GET("/oauth/user", oauth.GetUser)


	app.GET("/oauth/client/callback", Callback)

	app.GET("/oauth/client/login", func(context *gin.Context) {
		// Generate random state
		b := make([]byte, 32)
		_, err := rand.Read(b)
		if err != nil {
			 context.String(http.StatusOK, err.Error())
		}

		state := base64.StdEncoding.EncodeToString(b)

		stateStore[state] = true

		// Will return loginURL,
		// for example: http://localhost:4444/oauth2/auth?client_id=myclient&prompt=consent&redirect_uri=http%3A%2F%2Fexample.com&response_type=code&scope=users.write+users.read+users.edit&state=XfFcFf7KL7ajzA2nBY%2F8%2FX3lVzZ6VZ0q7a8rM3kOfMM%3D
		loginURL := OAuthConf.AuthCodeURL(state)
		context.Redirect(http.StatusFound, loginURL)
	})


	app.GET("/test",middlewares.OauthRequired() ,func(c *gin.Context) {
		c.JSON(200,gin.H{
			"success": "adadadada",
		})
	})

	return app
}

func Callback(c *gin.Context) {

	code := c.Query("code")
	state := c.Query("state")
	//scopes := c.QueryParam("scope")

	if code == "" {
		c.String(http.StatusOK, "authorization code is empty")
	}

	// If state is exist
	if _, exist := stateStore[state]; !exist {
		c.String(http.StatusOK, "state is generated by this Client")
	}

	delete(stateStore, state)

	// Exchange code for access token
	accessToken, err := OAuthConf.Exchange(c, code)
	if err != nil {
		 c.String(http.StatusOK, err.Error())
	}

	 c.JSON(http.StatusOK, gin.H{
		"AccessToken": accessToken,
	})
}



