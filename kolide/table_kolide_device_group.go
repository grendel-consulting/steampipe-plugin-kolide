package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideDeviceGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_device_group",
		Description: "Group of devices, created in the UI or API",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for this group.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "Human-readable name for this group.", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "When this group record was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "description", Description: "Longer-form description about this device group.", Type: proto.ColumnType_STRING},
			// Other columns
			{Name: "member_count", Description: "Number of member devices in this group.", Type: proto.ColumnType_INT},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this group.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide API query feature, can be combined with AND (and OR)
				{Name: "name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "created_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "description", Require: plugin.Optional, Operators: []string{"=", "~~"}},
			},
			Hydrate: listDeviceGroups,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDeviceGroup,
		},
	}
}

//// LIST FUNCTION

func listDeviceGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetDeviceGroups(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_device_group.listDeviceGroups", visitor, "DeviceGroups")
}

//// GET FUNCTION

func getDeviceGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetDeviceGroupById(id)
	}

	return getAnything(ctx, d, h, "kolide_device_group.getDeviceGroup", "id", visitor)
}
