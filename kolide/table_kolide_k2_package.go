package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideK2Package(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_k2_package",
		Description: "Installation packages for Kolide Launcher agent for each major operating system.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Unique identifier for this package.", Type: proto.ColumnType_STRING},
			// Other columns
			{Name: "built_at", Description: "When this installation package was built.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "url", Description: "URL that can be used to download this installation package. Requests to this url require the standard Authorization header needed for all Kolide K2 API requests.", Type: proto.ColumnType_STRING},
			{Name: "version", Description: "Version of the Launcher agent that will be installed by this package.", Type: proto.ColumnType_STRING},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this event.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Id")},
		},
		List: &plugin.ListConfig{
			Hydrate: listPackages,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getPackage,
		},
	}
}

//// LIST FUNCTION

func listPackages(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetPackages(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_k2_audit_log.listPackages", visitor, "Packages")
}

//// GET FUNCTION

func getPackage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetPackageById(id)
	}

	return getAnything(ctx, d, h, "kolide_k2_audit_log.getPackage", "id", visitor)
}
