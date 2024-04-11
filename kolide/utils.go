package kolide

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"

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
		return nil, errors.New("kolide 'api_token' must be set in the connection configuration; edit your connection configuration file and then restart")
	}

	c := kolide.New(
		kolide.WithAuth(api_token),
	)

	d.ConnectionManager.Cache.Set(cacheKey, c)

	return c, nil
}

type ListPredicate func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error)
type GetPredicate func(client *kolide.Client, id string) (interface{}, error)
type ListByIdPredicate func(client *kolide.Client, id string, cursor string, limit int32, searches ...kolide.Search) (interface{}, error)

func listAnything(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, callee string, visitor ListPredicate, target string) (interface{}, error) {
	// Create a slice to hold search queries
	searches, err := query(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error(callee, "qualifier_operator_error", err)
		return nil, err
	}

	// Establish connection to Kolide client
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error(callee, "connection_error", err)
		return nil, err
	}

	// Iterate through pagination cursors, with smallest number of pages
	var maxLimit int32 = kolide.MaxPaging
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	cursor := ""

	for {
		// Respect rate limiting
		d.WaitForListRateLimit(ctx)

		res, err := visitor(client, cursor, maxLimit, searches...)
		if err != nil {
			plugin.Logger(ctx).Error(callee, err)
			return nil, err
		}

		// Stream retrieved results
		collection := reflect.ValueOf(res).Elem().FieldByName(target)
		if collection.IsValid() {
			for i := 0; i < collection.Len(); i++ {
				d.StreamListItem(ctx, collection.Index(i).Interface())

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		}

		next := reflect.ValueOf(res).Elem().FieldByName("Pagination").FieldByName("NextCursor")
		if next.IsValid() {
			cursor = next.Interface().(string)
		}

		if cursor == "" {
			break
		}
	}

	return nil, nil
}

func getAnything(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, callee string, id string, visitor GetPredicate) (interface{}, error) {
	// Fail early if unique identifier is not present
	uid := d.EqualsQualString(id)
	if uid == "" {
		return nil, nil
	}

	// Establish connection to Kolide client
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error(callee, "connection_error", err)
		return nil, err
	}

	// Retrieve device based on id
	res, err := visitor(client, uid)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func listAnythingById(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, callee string, id string, visitor ListByIdPredicate, target string) (interface{}, error) {
	// Fail early if unique identifier is not present
	uid := d.EqualsQualString(id)
	if uid == "" {
		return nil, nil
	}
	// Create a slice to hold search queries
	searches, err := query(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error(callee, "qualifier_operator_error", err)
		return nil, err
	}

	// Establish connection to Kolide client
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error(callee, "connection_error", err)
		return nil, err
	}

	// Iterate through pagination cursors, with smallest number of pages
	var maxLimit int32 = kolide.MaxPaging
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	cursor := ""

	for {
		// Respect rate limiting
		d.WaitForListRateLimit(ctx)

		res, err := visitor(client, uid, cursor, maxLimit, searches...)
		if err != nil {
			plugin.Logger(ctx).Error(callee, err)
			return nil, err
		}

		// Stream retrieved results
		collection := reflect.ValueOf(res).Elem().FieldByName(target)
		if collection.IsValid() {
			for i := 0; i < collection.Len(); i++ {
				d.StreamListItem(ctx, collection.Index(i).Interface())

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		}

		next := reflect.ValueOf(res).Elem().FieldByName("Pagination").FieldByName("NextCursor")
		if next.IsValid() {
			cursor = next.Interface().(string)
		}

		if cursor == "" {
			break
		}
	}

	return nil, nil
}

func query(ctx context.Context, d *plugin.QueryData) ([]kolide.Search, error) {
	// Create a slice to hold search queries
	searches := make([]kolide.Search, 0)

	// Extract search queries from qualifiers
	for _, item := range d.Quals {
		for _, q := range item.Quals {
			search, err := mapToSearch(item.Name, q.Operator, q.Value.GetStringValue())
			if err != nil {
				return nil, err
			}
			searches = append(searches, search)
		}
	}

	return searches, nil
}

// Operators supported by Kolide API are:
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
