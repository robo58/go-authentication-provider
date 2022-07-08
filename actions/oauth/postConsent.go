package oauth

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go"
	"github.com/robo58/go-authentication-provider/data"
	"github.com/robo58/go-authentication-provider/data/models/users"
	"net/http"
)

type ConsentFormData struct {
	ConsentChallenge string   `validate:"required" form:"consent_challenge"`
	GrantScope       []string `validate:"required" form:"grant_scope"`
}

func PostConsent(c *gin.Context){
	var formData ConsentFormData

	bindErr := c.ShouldBind(&formData)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to bind form data",
		})
	}
	var user users.User
	session := sessions.Default(c)
	data.GetDB().First(&user, session.Get("user"))
	admin := client.NewAPIClient(data.GetHydraConfig().Admin).AdminApi

	consentGetParams := admin.GetConsentRequest(c).ConsentChallenge(formData.ConsentChallenge)

	consentGetResp, _, getErr := consentGetParams.Execute()
	if getErr != nil {
		// if error, redirects to ...
		str := fmt.Sprint("error GetConsentRequest", getErr.Error())
		c.String(http.StatusUnprocessableEntity, str)
	}

	// If a user has granted this application the requested scope, hydra will tell us to not show the UI.
	remember := false
	if session.Get("remember") != nil {
		remember = session.Get("remember").(bool)
	}
	var rememberFor int64
	rememberFor=30
	// Now it's time to grant the consent request. You could also deny the request if something went terribly wrong
	consentAcceptBody := client.AcceptConsentRequest{
		GrantAccessTokenAudience: consentGetResp.GetRequestedAccessTokenAudience(),
		GrantScope:               formData.GrantScope,
		Remember:                 &remember,
		RememberFor:              &rememberFor,
		Session: &client.ConsentRequestSession{
			// Sets session data for the OpenID Connect ID token.
			IdToken: map[string]interface{}{
				"user": map[string]interface{}{
					"id":        user.ID,
					"email": user.Email,
					"name":  user.Name,
				},
			},
		},
	}

	consentAcceptParams := admin.AcceptConsentRequest(c).ConsentChallenge(formData.ConsentChallenge).AcceptConsentRequest(consentAcceptBody)
	consentAcceptResp, _, acceptErr := consentAcceptParams.Execute()
	if acceptErr != nil {
		str := fmt.Sprint("error AcceptConsentRequest", acceptErr.Error())
		c.String(http.StatusUnprocessableEntity, str)
	}

	 c.Redirect(http.StatusFound, consentAcceptResp.GetRedirectTo())
}
