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
			{Name: "name", Description: "Author-supplied title of this live query campaign.", Type: proto.ColumnType_STRING},
			{Name: "osquery_sql", Description: "SQL query to be executed against all targeted devices.", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "When this live query campaign was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "published", Description: "Describes whether or not the author has published this live query campaign, or it is in draft.", Type: proto.ColumnType_BOOL},
			{Name: "revision", Description: "Editing a Live Query will increment its revision. When the revision is incremented, the live query will be re-run on all target devices.", Type: proto.ColumnType_INT},
			{Name: "tables_used", Description: "List of tables referenced in the SQL.", Type: proto.ColumnType_JSON},
			{Name: "successful_devices_count", Description: "Devices that have successfully run the query.", Type: proto.ColumnType_INT},
			{Name: "errored_devices_count", Description: "Devices that returned an error when attempting to run the query.", Type: proto.ColumnType_INT},
			{Name: "waiting_devices_count", Description: "Devices that have not yet reported results or errors running the device.", Type: proto.ColumnType_INT},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this live query campaign.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "created_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "published", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "tables_used", Require: plugin.Optional, Operators: []string{"=", "~~"}},
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
