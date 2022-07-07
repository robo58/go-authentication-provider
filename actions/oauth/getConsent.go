package oauth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go"
	"github.com/robo58/go-authentication-provider/data"
	"net/http"
	"strings"
)

func GetConsent(c *gin.Context)  {

	consentChallenge := strings.TrimSpace(c.Query("consent_challenge"))
	if consentChallenge == "" {
		 c.HTML(http.StatusOK, "consent.tmpl", gin.H{
			"ErrorTitle":   "Cannot Accept Consent Request",
			"ErrorContent": "Consent challenge is empty",
		})
	}

	admin := client.NewAPIClient(data.GetHydraConfig().Admin).AdminApi


	consentGetParams := admin.GetConsentRequest(c).ConsentChallenge(consentChallenge)

	consentGetResp, _, getErr := consentGetParams.Execute()
	if getErr != nil {
		c.HTML(http.StatusOK, "consent.tmpl", gin.H{
			"ErrorTitle":   "Cannot Accept Consent Request",
			"ErrorContent": getErr.Error(),
		})
	}


	// If a user has granted this application the requested scope, hydra will tell us to not show the UI.
	if consentGetResp.GetSkip() {
		// You can apply logic here, for example grant another scope, or do whatever...
		// ...

		// Now it's time to grant the consent request.
		// You could also deny the request if something went terribly wrong
		consentAcceptBody := client.AcceptConsentRequest{
			GrantAccessTokenAudience: consentGetResp.GetRequestedAccessTokenAudience(),
			GrantScope:               consentGetResp.GetRequestedScope(),
		}

		consentAcceptParams := admin.AcceptConsentRequest(c).ConsentChallenge(consentChallenge).AcceptConsentRequest(consentAcceptBody)

		consentAcceptResp, _, acceptErr := consentAcceptParams.Execute()
		if acceptErr != nil {
			str := fmt.Sprint("error AcceptConsentRequest", acceptErr.Error())
			c.String(http.StatusUnprocessableEntity, str)

		}

		 c.Redirect(http.StatusFound, consentAcceptResp.GetRedirectTo())
	}

	consentMessage := fmt.Sprintf("Application %s wants access resources on your behalf and to:",
		consentGetResp.GetClient().ClientName,
	)

	 c.HTML(http.StatusOK, "consent.tmpl", gin.H{
		"ConsentChallenge": consentChallenge,
		"ConsentMessage":   consentMessage,
		"RequestedScopes":  consentGetResp.GetRequestedScope(),
	})
}
