package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/mildred/terraform-provider-systemd/systemd"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: systemd.Provider,
	})
}
