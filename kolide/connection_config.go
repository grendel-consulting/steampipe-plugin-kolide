package kolide

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type kolideK2Config struct {
	APIToken *string `cty:"api_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &kolideK2Config{}
}

func GetConfig(connection *plugin.Connection) kolideK2Config {
	if connection == nil || connection.Config == nil {
		return kolideK2Config{}
	}
	config, _ := connection.Config.(kolideK2Config)
	return config
}
