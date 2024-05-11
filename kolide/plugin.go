package kolide

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-kolide",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromCamel().NullIfEmptySlice().NullIfZero(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		DefaultRetryConfig: &plugin.RetryConfig{
			// Preference would be to respect `Retry-After` header that Kolide API supports
			// For now, default to Fibonacci, ten retries starting at 100ms
			ShouldRetryErrorFunc: shouldRetryError([]string{"429"}),
		},

		TableMap: map[string]*plugin.Table{
			"kolide_admin_user":               tableKolideAdminUser(ctx),
			"kolide_audit_log":                tableKolideAuditLog(ctx),
			"kolide_check":                    tableKolideCheck(ctx),
			"kolide_deprovisioned_person":     tableKolideDeprovisionedPerson(ctx),
			"kolide_device":                   tableKolideDevice(ctx),
			"kolide_device_group":             tableKolideDeviceGroup(ctx),
			"kolide_device_group_device":      tableKolideDeviceGroupDevice(ctx),
			"kolide_device_open_issue":        tableKolideDeviceOpenIssue(ctx),
			"kolide_issue":                    tableKolideIssue(ctx),
			"kolide_package":                  tableKolidePackage(ctx),
			"kolide_person":                   tableKolidePerson(ctx),
			"kolide_person_group":             tableKolidePersonGroup(ctx),
			"kolide_person_open_issue":        tableKolidePersonOpenIssue(ctx),
			"kolide_person_registered_device": tableKolidePersonRegisteredDevice(ctx),
			"kolide_registration_request":     tableKolideRegistrationRequest(ctx),
		},
	}
	return p
}
