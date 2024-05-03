package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolidePersonGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_person_group",
		Description: "Group of people, these are synced from your SCIM provider",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for this group.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "Human-readable name for this group.", Type: proto.ColumnType_STRING},
			// Other columns
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this group.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide API query feature, can be combined with AND (and OR)
				{Name: "name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
			},
			Hydrate: listPersonGroups,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getPersonGroup,
		},
	}
}

//// LIST FUNCTION

func listPersonGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetPersonGroups(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_person_group.listPersonGroups", visitor, "PersonGroup")
}

//// GET FUNCTION

func getPersonGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetPersonGroupById(id)
	}

	return getAnything(ctx, d, h, "kolide_person_group.getPersonGroup", "id", visitor)
}
