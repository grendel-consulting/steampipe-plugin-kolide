package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideK2DeprovisionedPerson(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_k2_deprovisioned_person",
		Description: "Anyone who has been removed from Kolide via SCIM.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "name", Description: "Canonical human name for this person.", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "Recorded email address for this person.", Type: proto.ColumnType_STRING},
			{Name: "last_authenticated_at", Description: "When the person was last authenticated with Kolide.", Type: proto.ColumnType_TIMESTAMP},
			// Other columns
			{Name: "id", Description: "Canonical identifier for this person.", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "When the person record was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "has_registered_device", Description: "Whether or not this person has at least one registered device.", Type: proto.ColumnType_BOOL},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this person.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide K2 API query feature, can be combined with AND (and OR)
				{Name: "name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "email", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "last_authenticated_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
			},
			Hydrate: listDeprovisionedPeople,
		},
	}
}

//// LIST FUNCTION

func listDeprovisionedPeople(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetDeprovisionedPeople(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_k2_deprovisioned_person.listDeprovisionedPeople", visitor, "DeprovisionedPeople")
}
