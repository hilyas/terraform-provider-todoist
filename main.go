package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hilyas/terraform-provider-todoist/todoist"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: todoist.Provider,
	})
}
