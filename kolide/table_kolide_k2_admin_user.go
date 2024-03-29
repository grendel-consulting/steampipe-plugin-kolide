package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideK2AdminUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_k2_admin_user",
		Description: "Users with access to the Kolide dashboard.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for this admin user.", Type: proto.ColumnType_STRING},
			{Name: "first_name", Description: "First name for this admin user.", Type: proto.ColumnType_STRING},
			{Name: "last_name", Description: "Last human name for this admin user.", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "Email address for this admin user.", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "When this admin user account was created.", Type: proto.ColumnType_TIMESTAMP},
			// Other columns
			{Name: "access", Description: "Access level granted to this admin user, one of full, limited or billing.", Type: proto.ColumnType_STRING},
			{Name: "restrictions", Description: "Feature restrictions applied to this user; this list will be empty unless the user has an access level of 'limited'.", Type: proto.ColumnType_JSON},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this admin user.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide K2 API query feature, can be combined with AND (and OR)
				{Name: "first_name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "last_name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "email", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "created_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
			},
			Hydrate: listAdminUsers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAdminUser,
		},
	}
}

//// LIST FUNCTION

func listAdminUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetAdminUsers(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_k2_admin_user.listAdminUsers", visitor, "AdminUsers")
}

//// GET FUNCTION

func getAdminUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetAdminUserById(id)
	}

	return getAnything(ctx, d, h, "kolide_k2_admin_user.getAdminUser", "id", visitor)
}
