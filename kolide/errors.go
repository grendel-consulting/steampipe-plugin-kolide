package kolide

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func shouldIgnoreErrors(ignorableErros []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		for _, pattern := range ignorableErros {
			if strings.Contains(err.Error(), pattern) {
				return true
			}
		}
		return false
	}
}

func shouldRetryError(retryErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		for _, pattern := range retryErrors {
			if strings.Contains(err.Error(), pattern) {
				plugin.Logger(ctx).Debug("kolide_errors.shouldRetryError", "rate_limit_error", err)
				return true
			}
		}
		return false
	}
}
