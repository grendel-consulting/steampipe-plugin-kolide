package kolide

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type kolideConfig struct {
	APIToken *string `cty:"api_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &kolideConfig{}
}

func GetConfig(connection *plugin.Connection) kolideConfig {
	if connection == nil || connection.Config == nil {
		return kolideConfig{}
	}
	config, _ := connection.Config.(kolideConfig)
	return config
}
