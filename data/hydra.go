package data

import (
	client "github.com/ory/hydra-client-go"
	appConfig "github.com/robo58/go-authentication-provider/config"
)


type HydraConfig struct {
	Admin *client.Configuration
	Public *client.Configuration
}

var HydraInstance *HydraConfig

func SetupHydra() {
	hydraInstance := &HydraConfig{}
	config := appConfig.GetConfig()
	hydraInstance.Admin = client.NewConfiguration()
	hydraInstance.Admin.Servers = []client.ServerConfiguration{
		{
			URL: config.Hydra.AdminUrl, // Admin API URL
		},
	}
	hydraInstance.Public = client.NewConfiguration()
	hydraInstance.Public.Servers = []client.ServerConfiguration{
		{
			URL: config.Hydra.PublicUrl, // Public API URL
		},
	}

	HydraInstance = hydraInstance
}

func GetHydraConfig() *HydraConfig {
	return HydraInstance
}
