package kolide

import (
	"context"
	"errors"
	"fmt"
	"os"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	quals "github.com/turbot/steampipe-plugin-sdk/v5/plugin/quals"
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

// Operators supported by Kolide K2 API are:
//
// - Exact match, ":"
// - Substring match, "~", treated as Like
// - Greater than (for datetime fields only), ">"
// - Less than (for datetime fields only), "<"
//
// See: https://www.kolide.com/docs/developers/api#search
var operatorMapping = map[string]kolide.OperatorType{
	quals.QualOperatorEqual:   kolide.Equals,
	quals.QualOperatorLike:    kolide.SubstringMatch,
	quals.QualOperatorGreater: kolide.GreaterThan,
	quals.QualOperatorLess:    kolide.LessThan,
}

func mapToSearch(field string, qualifier string, value string) (kolide.Search, error) {
	if _, ok := operatorMapping[qualifier]; !ok {
		return kolide.Search{}, fmt.Errorf("unsupported qualifier: %s", qualifier)
	}
	return kolide.Search{
		Field:    field,
		Operator: operatorMapping[qualifier],
		Value:    value,
	}, nil
}
