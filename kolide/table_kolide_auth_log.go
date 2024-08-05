package kolide

import (
	"context"

	kolide "github.com/grendel-consulting/steampipe-plugin-kolide/kolide/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableKolideAuthLog(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kolide_auth_log",
		Description: "Authentication events occurring when a user tries to sign into an App protected by Kolide Device Trust.",
		Columns: []*plugin.Column{
			// Filterable "top" columns
			{Name: "id", Description: "Canonical identifier for this auth event.", Type: proto.ColumnType_STRING},
			{Name: "timestamp", Description: "When this event started.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "ip_address", Description: "IP address of the request intiating this auth event, may be IPv4 or IPv6.", Type: proto.ColumnType_STRING},
			{Name: "agent_version", Description: "Version of the Kolide Agent running on the endpoint, if known.", Type: proto.ColumnType_STRING},
			{Name: "country", Description: "Name of the country that the session originated from, determined by IP addres and subject to the limitations of IP geocoding.", Type: proto.ColumnType_STRING},
			{Name: "city", Description: "Name of the city that the session originated from, determined by IP addres and subject to the limitations of IP geocoding.", Type: proto.ColumnType_STRING},
			{Name: "browser_name", Description: "Common name of the browser used to initiate the session, subject to the limitations and accuracy of browser detection.", Type: proto.ColumnType_STRING},
			// Other columns
			{Name: "person_name", Description: "Name of the user triggering this auth event.", Type: proto.ColumnType_STRING},
			{Name: "person_email", Description: "Email of the user triggering this auth event.", Type: proto.ColumnType_STRING},
			{Name: "person_id", Description: "Canonical identifier for the user this auth event relates to.", Type: proto.ColumnType_STRING, Transform: transform.FromField("PersonInformation.Identifier")},
			{Name: "device_id", Description: "Canonical identifier for the device this auth event relates to.", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeviceInformation.Identifier")},
			{Name: "result", Description: "Result of the authentication attempt, either Success or Fail.", Type: proto.ColumnType_STRING},
			{Name: "initial_status", Description: "Initial auth status of the device attempting authentication, one of All_Good, Will_Block, Blocked or Unknown if no device was detected.", Type: proto.ColumnType_STRING},
			{Name: "browser_user_agent", Description: "User agent information for the browser used to initiatie this session, subject to the limitations and accuracy of browser detection.", Type: proto.ColumnType_STRING},
			{Name: "issues_displayed", Description: "List of issue titles and blocking status that were displayed to the end user", Type: proto.ColumnType_JSON, Transform: transform.FromField("IssuesDisplayed")},
			{Name: "events", Description: "Events that occured during this authentication session", Type: proto.ColumnType_JSON, Transform: transform.FromField("Events")},
			// Steampipe standard columns
			{Name: "title", Description: "Display name for this event.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Result")},
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				// Using Kolide API query feature, can be combined with AND (and OR)
				{Name: "timestamp", Require: plugin.Optional, Operators: []string{"=", ">", "<"}},
				{Name: "city", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "country", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "ip_address", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "agent_version", Require: plugin.Optional, Operators: []string{"=", "~~"}},
				{Name: "browser_name", Require: plugin.Optional, Operators: []string{"=", "~~"}},
			},
			Hydrate: listAuthLogs,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAuthLog,
		},
	}
}

//// LIST FUNCTION

func listAuthLogs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor ListPredicate = func(client *kolide.Client, cursor string, limit int32, searches ...kolide.Search) (interface{}, error) {
		return client.GetAuthLogs(cursor, limit, searches...)
	}

	return listAnything(ctx, d, h, "kolide_auth_log.listAuthLogs", visitor, "AuthLogs")
}

//// GET FUNCTION

func getAuthLog(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var visitor GetPredicate = func(client *kolide.Client, id string) (interface{}, error) {
		return client.GetAuthLogById(id)
	}

	return getAnything(ctx, d, h, "kolide_auth_log.getAuthLog", "id", visitor)
}
