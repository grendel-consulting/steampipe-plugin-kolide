package main

import (
	"github.com/grendel-consulting/steampipe-plugin-kolide/kolide"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: kolide.Plugin})
}
