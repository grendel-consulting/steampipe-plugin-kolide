package kolide

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideK2Device(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_k2_device",
		Description: "Devices enrolled and monitored by Kolide.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for the device.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "Canonical human name for the device.", Type: proto.ColumnType_STRING},
			{Name: "registered_at", Description: "When the device was registered to its current owner.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "last_authenticated_at", Description: "When the device was last authenticated with Kolide.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "serial", Description: "Hardware serial for the device.", Type: proto.ColumnType_STRING},
			{Name: "hardware_uuid", Description: "Hardware UUID for the device.", Type: proto.ColumnType_STRING},
			{Name: "note", Description: "Notes provided by a Kolide administrator (in Markdown format).", Type: proto.ColumnType_STRING},
			{Name: "will_block_at", Description: "If the auth status is 'Will Block', this timestampe describes when the device will be blocked by a failing check.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "device_type", Description: "Platform type of the device, one of Mac, Windows, Linux, iOS or Android.", Type: proto.ColumnType_STRING},
			// Other columns
			// RegisteredOwnerInfo {}
			{Name: "operating_system", Description: "Operating system installed on the device.", Type: proto.ColumnType_STRING},
			{Name: "hardware_model", Description: "Specific hardware model of the device.", Type: proto.ColumnType_STRING},
			{Name: "auth_state", Description: "Authorisation status of the device, one of Good, Notified, Will Block or Blocked.", Type: proto.ColumnType_STRING},
			{Name: "product_image_url", Description: "URL of the device's product image.", Type: proto.ColumnType_STRING},
			// AuthConfiguration {}
			{Name: "form_factor", Description: "Form factor of the device, one of Computer, Tablet or Phone.", Type: proto.ColumnType_STRING},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this resource.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide K2 API query feature, can be combined with AND (and OR)
				{Name: "name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "registered_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "last_authenticated_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "serial", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "note", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "hardware_uuid", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "device_type", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "will_block_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
			},
			Hydrate: listDevices,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDevice,
		},
	}
}

//// LIST FUNCTION

func listDevices(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create a slice to hold search queries
	searches, err := query(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("kolide_k2_device.listDevices", "qualifier_operator_error", err)
		return nil, err
	}

	// Establish connection to Kolide client
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("kolide_k2_device.listDevices", "connection_error", err)
		return nil, err
	}

	// Iterate through pagination cursors
	cursor := ""

	for {
		// Respect rate limiting
		d.WaitForListRateLimit(ctx)

		res, err := client.GetDevices(cursor, searches...)
		if err != nil {
			plugin.Logger(ctx).Error("kolide_k2_device.listDevices", err)
			return nil, err
		}

		// Stream retrieved devices
		for _, device := range res.Devices {
			d.StreamListItem(ctx, device)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		cursor = res.Pagination.NextCursor

		if cursor == "" {
			break
		}
	}

	return nil, nil
}

//// GET FUNCTION

func getDevice(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Fail early if "id" is not present
	id := d.EqualsQualString("id")
	if id == "" {
		return nil, nil
	}

	// Establish connection to Kolide client
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("kolide_k2_device.getDevice", "connection_error", err)
		return nil, err
	}

	// Retrieve device based on id
	res, err := client.GetDeviceById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
