package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robo58/go-authentication-provider/middlewares"
	"io"
	"net/http"
	"os"
)

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

	//// Routes
	//// ================== Login Routes
	////app.POST("/api/login", controllers.Login)
	//app.POST("/api/register", controllers.CreateUser)
	//// ================== Docs Routes
	////app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//// ================== User Routes
	//app.GET("/api/users", controllers.GetUsers)
	//app.GET("/api/users/:id", controllers.GetUserById)
	//app.POST("/api/users", controllers.CreateUser)
	//app.PUT("/api/users/:id", controllers.UpdateUser)
	//app.DELETE("/api/users/:id", controllers.DeleteUser)

	app.GET("/api/test", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"message": "OK",
		})
	})

	return app
}
