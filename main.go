package main

import (
	"github.com/hashicorp/terraform/plugin"
	"terraform-provider-littlewarden/littlewarden"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: littlewarden.Provider,
	})
}
