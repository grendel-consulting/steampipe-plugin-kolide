package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
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
			// {
			// 	Name:        "akas",
			// 	Description: "Array of also-known-as identifiers that uniquely and immutably identify this resource.",
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
				// Using Kolide K2 API path, uniquely filters
				{Name: "id", Require: plugin.Optional, Operators: []string{"="}},
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
	if d.Quals["id"] != nil {
		plugin.Logger(ctx).Debug("kolide_k2_device.listDevices", "short-circuit by id")
		return getDevice(ctx, d, h)
	}

	var searches []kolide.Search

	if len(d.Quals) > 0 {
		for _, item := range d.Quals {
			for _, q := range d.Quals[item.Name].Quals {
				search, err := mapToSearch(item.Name, q.Operator, q.Value.GetStringValue())

				if err != nil {
					plugin.Logger(ctx).Error("kolide_k2_device.listDevices", "qualifier_operator_error", err)
					return nil, err
				}

				searches = append(searches, search)
			}
		}
	}

	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("kolide_k2_device.listDevices", "connection_error", err)
		return nil, err
	}

	res, err := client.GetDevices(searches...)
	if err != nil {
		plugin.Logger(ctx).Error("kolide_k2_device.listDevices", err)
		return nil, err
	}
	for _, user := range res.Devices {
		d.StreamListItem(ctx, user)
	}
	return nil, nil
}

//// GET FUNCTION

func getDevice(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	if id == "" {
		return nil, nil
	}

	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("kolide_k2_device.getDevice", "connection_error", err)
		return nil, err
	}

	result, err := client.GetDeviceById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
