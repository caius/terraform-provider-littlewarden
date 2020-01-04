package main

import (
	"terraform-provider-littlewarden/littlewarden"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: littlewarden.Provider,
	})
}
