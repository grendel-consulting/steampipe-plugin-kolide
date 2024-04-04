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
			// Preference would be to respect Retry-After header that Kolide K2 API supports
			// For now, default to Fibonacci, ten retries starting at 100ms
			ShouldRetryErrorFunc: shouldRetryError([]string{"429"}),
		},

		TableMap: map[string]*plugin.Table{
			"kolide_k2_admin_user":           tableKolideK2AdminUser(ctx),
			"kolide_k2_audit_log":            tableKolideK2AuditLog(ctx),
			"kolide_k2_check":                tableKolideK2Check(ctx),
			"kolide_k2_deprovisioned_person": tableKolideK2DeprovisionedPerson(ctx),
			"kolide_k2_device":               tableKolideK2Device(ctx),
			"kolide_k2_device_open_issue":    tableKolideK2DeviceOpenIssue(ctx),
			"kolide_k2_issue":                tableKolideK2Issue(ctx),
			"kolide_k2_package":              tableKolideK2Package(ctx),
		},
	}
	return p
}
