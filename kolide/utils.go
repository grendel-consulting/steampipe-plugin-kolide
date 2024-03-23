package kolide

import (
	"context"
	"errors"
	"os"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*kolide.Client, error) {
	cacheKey := "kolide"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*kolide.Client), nil
	}

	kolideConfig := GetConfig(d.Connection)

	api_token := os.Getenv("KOLIDE_API_TOKEN")

	if kolideConfig.APIToken != nil {
		api_token = *kolideConfig.APIToken
	}

	if api_token == "" {
		return nil, errors.New("kolide k2 'api_token' must be set in the connection configuration; edit your connection configuration file and then restart")
	}

	c := kolide.New(
		kolide.WithAuth(api_token),
	)

	d.ConnectionManager.Cache.Set(cacheKey, c)

	return c, nil
}
