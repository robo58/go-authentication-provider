package oauth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go"
	"github.com/robo58/go-authentication-provider/data"
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

	admin := client.NewAPIClient(data.GetHydraConfig().Admin).AdminApi


	consentGetParams := admin.GetConsentRequest(c).ConsentChallenge(formData.ConsentChallenge)

	consentGetResp, _, getErr := consentGetParams.Execute()
	if getErr != nil {
		// if error, redirects to ...
		str := fmt.Sprint("error GetConsentRequest", getErr.Error())
		c.String(http.StatusUnprocessableEntity, str)
	}

	// If a user has granted this application the requested scope, hydra will tell us to not show the UI.

	// Now it's time to grant the consent request. You could also deny the request if something went terribly wrong
	consentAcceptBody := client.AcceptConsentRequest{
		GrantAccessTokenAudience: consentGetResp.GetRequestedAccessTokenAudience(),
		GrantScope:               formData.GrantScope,
	}

	consentAcceptParams := admin.AcceptConsentRequest(c).ConsentChallenge(formData.ConsentChallenge).AcceptConsentRequest(consentAcceptBody)
	consentAcceptResp, _, acceptErr := consentAcceptParams.Execute()
	if acceptErr != nil {
		str := fmt.Sprint("error AcceptConsentRequest", acceptErr.Error())
		c.String(http.StatusUnprocessableEntity, str)
	}

	 c.Redirect(http.StatusFound, consentAcceptResp.GetRedirectTo())
}
