package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolidePersonRegisteredDevice(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_person_registered_device",
		Description: "Registered devices belonging to a person.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "person_id", Description: "Canonical identifier of the registered owner of the device.", Type: proto.ColumnType_STRING, Transform: transform.FromQual("person_id")},
			{Name: "id", Description: "Canonical identifier for the device.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "Canonical human name for the device.", Type: proto.ColumnType_STRING},
			{Name: "registered_at", Description: "When the device was registered to its current owner.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "last_authenticated_at", Description: "When the device was last authenticated with Kolide.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "serial", Description: "Hardware serial for the device.", Type: proto.ColumnType_STRING},
			{Name: "hardware_uuid", Description: "Hardware UUID for the device.", Type: proto.ColumnType_STRING},
			{Name: "note", Description: "Notes provided by a Kolide administrator (in Markdown format).", Type: proto.ColumnType_STRING},
			{Name: "will_block_at", Description: "If the auth status is 'Will Block', this timestamp describes when the device will be blocked by a failing check.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "device_type", Description: "Platform type of the device, one of Mac, Windows, Linux, iOS or Android.", Type: proto.ColumnType_STRING},
			// Other columns
			{Name: "registered_owner_identifier", Description: "Canonical identifier for the registered owner of this device.", Type: proto.ColumnType_STRING, Transform: transform.FromField("RegisteredOwnerInfo.Identifier")},
			{Name: "operating_system", Description: "Operating system installed on the device.", Type: proto.ColumnType_STRING},
			{Name: "hardware_model", Description: "Specific hardware model of the device.", Type: proto.ColumnType_STRING},
			{Name: "auth_state", Description: "Authorisation status of the device, one of Good, Notified, Will Block or Blocked.", Type: proto.ColumnType_STRING},
			{Name: "product_image_url", Description: "URL of the device's product image.", Type: proto.ColumnType_STRING},
			{Name: "auth_configuration_device_id", Description: "Canonical identifier for this device, empty if it is not registered", Type: proto.ColumnType_STRING, Transform: transform.FromField("AuthConfiguration.DeviceId")},
			{Name: "auth_configuration_authentication_mode", Description: "Who can be authenticated with this device, one of 'only_registered_owner', 'only_registered_owner_or_group_members' or 'anyone'.", Type: proto.ColumnType_STRING, Transform: transform.FromField("AuthConfiguration.AuthenticationMode")},
			{Name: "auth_configuration_person_groups", Description: "Description of the groups allowed to authenticate with this device.", Type: proto.ColumnType_JSON, Transform: transform.FromField("AuthConfiguration.PersonGroups")},
			{Name: "form_factor", Description: "Form factor of the device, one of Computer, Tablet or Phone.", Type: proto.ColumnType_STRING},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this resource.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "person_id", Require: plugin.Required, Operators: []string{"="}},
				// Using Kolide API query feature, can be combined with AND (and OR)
				{Name: "name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "registered_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "last_authenticated_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "serial", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "note", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "hardware_uuid", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "device_type", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "will_block_at", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
			},
			Hydrate: listDeviceByPerson,
		},
	}
}

//// LIST FUNCTION

func listDeviceByPerson(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListByIdPredicate = func(client *kolide.Client, id string, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetDevicesByPerson(id, cursor, limit, searches...)
	}

	return listAnythingById(ctx, d, h, "kolide_person_registered_device.listDeviceByPerson", "person_id", visitor, "Device")
}
