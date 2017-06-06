package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-oneandone/oneandone"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: oneandone.Provider})
}
