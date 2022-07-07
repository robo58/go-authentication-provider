package oauth

import (
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go"
	"github.com/robo58/go-authentication-provider/data"
	"net/http"
	"strings"
)


func GetLogin(c *gin.Context)  {
	loginChallenge := strings.TrimSpace(c.Query("login_challenge"))
	if loginChallenge == "" {
		c.JSON(400, gin.H{
			"error": "Login Challenge Is Not Present!",
		})
	}

	hydra := client.NewAPIClient(data.GetHydraConfig().Admin)
	loginGetParam := hydra.AdminApi.GetLoginRequest(c).LoginChallenge(loginChallenge)
	exec, _, err := loginGetParam.ApiService.GetLoginRequestExecute(loginGetParam)
	if err != nil {
		c.JSON(400,gin.H{
			"error": err.Error(),
		})
	}

	if exec.GetSkip() {
		// Using Hydra Admin to accept login request!
		loginAcceptParam := hydra.AdminApi.AcceptLoginRequest(c).LoginChallenge(loginChallenge).AcceptLoginRequest(client.AcceptLoginRequest{Subject: exec.GetSubject()})
		execute, _, err := hydra.AdminApi.AcceptLoginRequestExecute(loginAcceptParam)
		if err != nil {
			c.HTML(http.StatusOK, "login.tmpl", gin.H{
				"ErrorTitle":   "Cannot Accept Login Request",
				"ErrorContent": err.Error(),
			})
		}
		// If success, it will redirect to consent page using handler GetConsent
		// It then show the consent form
		c.Redirect(http.StatusFound, execute.GetRedirectTo())
	}

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"LoginChallenge": loginChallenge,
	})
}
