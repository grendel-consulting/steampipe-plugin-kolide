package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideLiveQueryCampaign(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_live_query_campaign",
		Description: "Live query campaigns in Kolide.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for this live query campaign.", Type: proto.ColumnType_STRING},
			{Name: "query", Description: "The SQL query executed in this live query campaign.", Type: proto.ColumnType_STRING},
			{Name: "status", Description: "Status of the live query campaign.", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "When this live query campaign was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "updated_at", Description: "When this live query campaign was last updated.", Type: proto.ColumnType_TIMESTAMP},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this live query campaign.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Query")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "query", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "status", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "created_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "updated_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
			},
			Hydrate: listLiveQueryCampaigns,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getLiveQueryCampaign,
		},
	}
}

//// LIST FUNCTION

func listLiveQueryCampaigns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetLiveQueryCampaigns(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_live_query_campaign.listLiveQueryCampaigns", visitor, "LiveQueryCampaigns")
}

//// GET FUNCTION

func getLiveQueryCampaign(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetLiveQueryCampaignById(id)
	}

	return getAnything(ctx, d, h, "kolide_live_query_campaign.getLiveQueryCampaign", "id", visitor)
}
