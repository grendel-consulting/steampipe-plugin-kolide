package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideRegistrationRequest(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_registration_request",
		Description: "Created when someone requests admin approval to register a device with Kolide. Registration requests can be approved or denied by admins",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for this registration request.", Type: proto.ColumnType_STRING},
			{Name: "status", Description: "Current status of this registration request; one of: pending, approved or denied", Type: proto.ColumnType_STRING},
			{Name: "requester_message", Description: "Any message the person provided when requesting approval for this registration.", Type: proto.ColumnType_STRING},
			{Name: "requested_at", Description: "When this registration approval was requested.", Type: proto.ColumnType_TIMESTAMP},
			// Other columns
			{Name: "end_user_denial_note", Description: "Any explanation the admin provided for the requester when denying this request, for internal documentation. It is meant to be shown to the requester. It will be blank if the registration has not been denied.", Type: proto.ColumnType_STRING},
			{Name: "internal_denial_note", Description: "Any internal explanation the admin provided when denying this request, for internal documentation. It is not shown to the requester. It will be blank if the registration has not been denied.", Type: proto.ColumnType_STRING},
			{Name: "requester_id", Description: "Canonical identifier for the person who requested this registration.", Type: proto.ColumnType_STRING, Transform: transform.FromField("RequesterInformation.Identifier")},
			{Name: "device_id", Description: "Canonical identifier for the device the person is trying to register.", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeviceInformation.Identifier")},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this registration request.", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeviceInformation.Identifier")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide API query feature, can be combined with AND (and OR)
				{Name: "status", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "requested_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "requester_message", Require: plugin.Optional, Operators: []string{"=", "~~"}},
			},
			Hydrate: listRegistrationRequests,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getRegistrationRequest,
		},
	}
}

//// LIST FUNCTION

func listRegistrationRequests(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetRegistrationRequests(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_registration_request.listRegistrationRequests", visitor, "RegistrationRequests")
}

//// GET FUNCTION

func getRegistrationRequest(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetRegistrationRequestById(id)
	}

	return getAnything(ctx, d, h, "kolide_registration_request.getRegistrationRequest", "id", visitor)
}
