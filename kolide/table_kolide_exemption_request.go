package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideExemptionRequest(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_exemption_request",
		Description: "Created when when a user desires to permanently ignore a specific issue or set of issues on a single device. Exemption requests are created by users and can be approved or denied by admins.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for this exemption request.", Type: proto.ColumnType_STRING},
			{Name: "status", Description: "Current status of this exemption request; one of: open, withdrawn, approved or denied", Type: proto.ColumnType_STRING},
			{Name: "requester_message", Description: "Any message the person provided when asking for this exemption.", Type: proto.ColumnType_STRING},
			{Name: "requested_at", Description: "When this exemption was requested.", Type: proto.ColumnType_TIMESTAMP},
			// Other columns
			{Name: "denial_explanation", Description: "Any explanation the admin provided for the requester when denying this request, for internal documentation. It is meant to be shown to the requester. It will be blank if the exemption has not been denied.", Type: proto.ColumnType_STRING},
			{Name: "internal_explanation", Description: "Any internal explanation the admin provided when approving or denying this request, for internal documentation. It is not shown to the requester. It will be blank if the exemption has not been denied.", Type: proto.ColumnType_STRING},
			{Name: "requester_id", Description: "Canonical identifier for the person who requested this exemption.", Type: proto.ColumnType_STRING, Transform: transform.FromField("RequesterInformation.Identifier")},
			{Name: "device_id", Description: "Canonical identifier for the device the person is trying to register.", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeviceInformation.Identifier")},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this exemption request.", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeviceInformation.Identifier")},
			{Name: "issues", Description: "", Type: proto.ColumnType_JSON, Transform: transform.FromField("Issues.Identifier")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide API query feature, can be combined with AND (and OR)
				{Name: "status", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "requested_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "requester_message", Require: plugin.Optional, Operators: []string{"=", "~~"}},
			},
			Hydrate: listExemptionRequests,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getExemptionRequest,
		},
	}
}

//// LIST FUNCTION

func listExemptionRequests(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetExemptionRequests(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_exemption_request.listExemptionRequests", visitor, "ExemptionRequests")
}

//// GET FUNCTION

func getExemptionRequest(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetExemptionRequestById(id)
	}

	return getAnything(ctx, d, h, "kolide_exemption_request.getExemptionRequest", "id", visitor)
}
