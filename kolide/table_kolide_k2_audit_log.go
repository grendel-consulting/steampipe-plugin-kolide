package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideK2AuditLog(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_k2_audit_log",
		Description: "Tracked events occurring in the Kolide web console.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for this audit log event.", Type: proto.ColumnType_STRING},
			{Name: "timestamp", Description: "When this event occurred.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "actor_name", Description: "Name of the entity triggering this event.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "Description of the event that occurred.", Type: proto.ColumnType_STRING},
			// Other columns
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this event.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide K2 API query feature, can be combined with AND (and OR)
				{Name: "timestamp", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "actor_name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "description", Require: plugin.Optional, Operators: []string{"=", "~~"}},
			},
			Hydrate: listAuditLogs,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAuditLog,
		},
	}
}

//// LIST FUNCTION

func listAuditLogs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetAuditLogs(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_k2_audit_log.listAuditLogs", visitor, "AuditLogs")
}

//// GET FUNCTION

func getAuditLog(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetAuditLogById(id)
	}

	return getAnything(ctx, d, h, "kolide_k2_audit_log.getAuditLog", "id", visitor)
}
