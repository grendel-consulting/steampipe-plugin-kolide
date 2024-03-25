package kolide

import (
	"context"

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
			{Name: "id", Description: "Canonical identifier for this admin user.", Type: proto.ColumnType_STRING},
			{Name: "first_name", Description: "First name for this admin user.", Type: proto.ColumnType_STRING},
			{Name: "last_name", Description: "Last human name for this admin user.", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "Email address for this admin user.", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "When this admin user account was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "access", Description: "Access level granted to this admin user, one of full, limited or billing.", Type: proto.ColumnType_STRING},
			// Restrictions {}
			// Steampipe standard columns
			{
				Name: "title", Description: "Display name for this admin user.", Type: proto.ColumnType_STRING,
				Transform: transform.FromField("email"),
			},
			// {
			// 	Name:        "akas",
			// 	Description: "Array of also-known-as identifiers that uniquely identify this resource.",
			// 	Type:        proto.ColumnType_JSON,
			// 	Transform:   ...
			// },
			// {
			// 	Name:        "tags",
			// 	Description: "Any tags or labels on this resource.",
			// 	Type:        proto.ColumnType_JSON,
			// 	Transform:   ...
			// },
		},
		List: &plugin.ListConfig{
			Hydrate: listAdminUsers,
		},
	}
}

//// LIST FUNCTION

func listAdminUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("kolide_k2_admin_user.listAdminUsers", "connection_error", err)
		return nil, err
	}
	res, err := client.GetAdminUsers()
	if err != nil {
		plugin.Logger(ctx).Error("kolide_k2_admin_user.listAdminUsers", err)
		return nil, err
	}
	for _, user := range res.AdminUsers {
		d.StreamListItem(ctx, user)
	}
	return nil, nil
}
