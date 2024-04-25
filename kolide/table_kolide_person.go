package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolidePerson(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_person",
		Description: "People within your organisation, who may not necessarily have access to the Kolide dashboard (see kolide_admin_user). Devices may be registered to a person",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Unique identifier for this person.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "Human-readable name for this person.", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "Recorded email addresss for this person.", Type: proto.ColumnType_STRING},
			{Name: "last_authenticated_at", Description: "When this person last authenticated with Kolide.", Type: proto.ColumnType_TIMESTAMP},
			// Other columns
			{Name: "created_at", Description: "When this person record was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "has_registered_device", Description: "Whether or not this person has at least one reigstered device.", Type: proto.ColumnType_BOOL},
			{Name: "usernames", Description: "Any usernames imported from the SCIM provider associated with this person.", Type: proto.ColumnType_JSON},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this person.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide API query feature, can be combined with AND (and OR)
				{Name: "name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "email", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "last_authenticated_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
			},
			Hydrate: listPeople,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getPerson,
		},
	}
}

//// LIST FUNCTION

func listPeople(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetPackages(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_person.listPeople", visitor, "Person")
}

//// GET FUNCTION

func getPerson(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetPersonById(id)
	}

	return getAnything(ctx, d, h, "kolide_person.getPerson", "id", visitor)
}
