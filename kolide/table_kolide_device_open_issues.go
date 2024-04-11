package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideDeviceOpenIssue(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_device_open_issue",
		Description: "Unresolved and non-exempt issues created when a specific device fails a check; some checks, when they fail, will produce multiple Issues, each with a unique primary_key_value.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for this issue.", Type: proto.ColumnType_STRING},
			{Name: "issue_key", Description: "Primary key that distinguishes one issue from another in the context of a single check; only applicable for checks that can produce multiple issues.", Type: proto.ColumnType_STRING},
			{Name: "issue_value", Description: "Primary identifying value that distinguishes one issue from another in the context of a single check; only applicable for checks that can produce multiple issues.", Type: proto.ColumnType_STRING},
			{Name: "title", Description: "Descriptive title for this issue.", Type: proto.ColumnType_STRING},
			{Name: "exempted", Description: "Whether this issue has been granted an exemption.", Type: proto.ColumnType_BOOL},
			{Name: "resolved_at", Description: "When this issue was resolved, or null if still open.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "detected_at", Description: "When this issue was initially detected.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "blocks_device_at", Description: "When the device will be blocked from authenticating by this failing issue, or null if the check is not configured to block authentication.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "device_id", Description: "Canonical identifier for the device this issue relates to.", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeviceInformation.Identifier")},
			{Name: "check_id", Description: "Canonical identifier for the check this issue relates to.", Type: proto.ColumnType_STRING, Transform: transform.FromField("CheckInformation.Identifier")},
			{Name: "last_rechecked_at", Description: "When this issue was last rechecked.", Type: proto.ColumnType_TIMESTAMP},
			// Other columns
			{Name: "value", Description: "Relevant data that describes why the device failed the check.", Type: proto.ColumnType_JSON},
			// Steampipe standard columns
			// - We include "title" above as an expected Kolide API column, and it is sufficient
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "device_id", Require: plugin.Required, Operators: []string{"="}},
				// Using Kolide API query feature, can be combined with AND (and OR)
				{Name: "issue_key", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "issue_value", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "title", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "exempted", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "resolved_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "detected_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "blocks_device_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "check_id", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "last_rechecked_at", Require: plugin.Optional, Operators: []string{"=", "~~"}},
			},
			Hydrate: listIssuesByDevice,
		},
	}
}

//// LIST FUNCTION

func listIssuesByDevice(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListByIdPredicate = func(client *kolide.Client, id string, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetIssuesByDevice(id, cursor, limit, searches...)
	}

	return listAnythingById(ctx, d, h, "kolide_device_open_issue.listIssuesByDevice", "device_id", visitor, "Issues")
}
