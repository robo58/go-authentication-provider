package oauth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go"
	db "github.com/robo58/go-authentication-provider/data"
	"github.com/robo58/go-authentication-provider/data/models/users"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginFormData struct {
	LoginChallenge string `validate:"required" form:"login_challenge"`
	Email          string `validate:"required" form:"email"`
	Password       string `validate:"required" form:"password"`
	RememberMe     string `validate:"required" form:"remember_me"`
}

func PostLogin(c *gin.Context)  {
	var formData LoginFormData
	err := c.ShouldBind(&formData)
	if err != nil {
		c.JSON(400,gin.H{
			"error": "Failed to bind form data",
		})
	}

	// TODO validation

	var rememberMe = formData.RememberMe == "true"
    // get user
	data := db.GetDB()
	var user users.User

	// check mail and password
	data.Where(&users.User{Email: formData.Email}).First(&user)

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))
	if passErr != nil {
		return
	}



	hydra := client.NewAPIClient(db.GetHydraConfig().Admin)

	// Using Hydra Admin to accept login request!
	loginGetParam := hydra.AdminApi.GetLoginRequest(c).LoginChallenge(formData.LoginChallenge)

	_, _, execError := loginGetParam.Execute()
	if execError != nil {
		// if error, redirects to ...
		str := fmt.Sprint("error GetLoginRequest", execError.Error())
		c.String(http.StatusUnprocessableEntity, str)
	}

	subject := fmt.Sprint(user.ID)

	loginAcceptParam := hydra.AdminApi.AcceptLoginRequest(c).LoginChallenge(formData.LoginChallenge).AcceptLoginRequest(client.AcceptLoginRequest{Subject: subject, Remember: &rememberMe})

	execute, _, acceptErr := loginAcceptParam.Execute()
	if acceptErr != nil {
		// if error, redirects to ...
		str := fmt.Sprint("error AcceptLoginRequest", acceptErr)
		c.String(http.StatusUnprocessableEntity, str)
	}


	// If success, it will redirect to consent page using handler GetConsent
	// It then show the consent form
	c.Redirect(http.StatusFound, execute.GetRedirectTo())
}