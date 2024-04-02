package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideK2Check(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_k2_check",
		Description: "Checks that Kolide runs on a device on a regular cadence, which are tests that typically produces a passing or failing result.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for this check.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "Descriptive name of the state this check is meant to enforce.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "Longer-form description of the check's purpose and operation.", Type: proto.ColumnType_STRING},
			{Name: "check_tags", Description: "List of Team-set tags associated with the check.", Type: proto.ColumnType_JSON},
			// Other columns
			{Name: "compatible_platforms", Description: "Array of device platforms this check can run on, taken from linux, windows or darwin.", Type: proto.ColumnType_JSON},
			{Name: "topics", Description: "List of Kolide-set topics associated with the check.", Type: proto.ColumnType_JSON},
			{Name: "blocking_enabled", Description: "Whether or not an issue for this check will block device trust authentication.", Type: proto.ColumnType_BOOL, Transform: transform.FromField("BlocksAuthConfiguration.BlockingEnabled")},
			{Name: "grace_period_days", Description: "Number of days that a device is allowed to be failing the check before it will be blocked.", Type: proto.ColumnType_INT, Transform: transform.FromField("BlocksAuthConfiguration.GracePeriodDays")},
			{Name: "blocking_group_names", Description: "List of names for the groups whose device members will be blocked by this check, unless excluded.", Type: proto.ColumnType_JSON, Transform: transform.FromField("BlocksAuthConfiguration.BlockingGroupNames")},
			{Name: "blocking_exempt_group_names", Description: "List of names for the groups whose device members are exemnpt from being blocked by this check.", Type: proto.ColumnType_JSON, Transform: transform.FromField("BlocksAuthConfiguration.BlockingExemptGroupNames")},
			{Name: "excluded_groups", Description: "List of names for the groups whose device members are excluded from being targets for this check.", Type: proto.ColumnType_JSON, Transform: transform.FromField("TargetingConfiguration.ExcludedGroups")},
			{Name: "targeted_groups", Description: "List of names for the groups whose device members will be targets for this check, unless excluded.", Type: proto.ColumnType_JSON, Transform: transform.FromField("TargetingConfiguration.TargetedGroups")},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this check.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide K2 API query feature, can be combined with AND (and OR)
				{Name: "name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "description", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "check_tags", Require: plugin.Optional, Operators: []string{"=", "~~"}},
			},
			Hydrate: listChecks,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getCheck,
		},
	}
}

//// LIST FUNCTION

func listChecks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetChecks(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_k2_audit_log.listChecks", visitor, "Checks")
}

//// GET FUNCTION

func getCheck(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetCheckById(id)
	}

	return getAnything(ctx, d, h, "kolide_k2_audit_log.getCheck", "id", visitor)
}
