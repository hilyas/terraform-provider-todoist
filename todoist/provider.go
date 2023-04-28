package todoist

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TODOIST_API_KEY", nil),
				Description: "API key for the Todoist API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"todoist_project": ResourceProject(),
			"todoist_task":    ResourceTask(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"todoist_project": DataSourceProject(),
			"todoist_task":    DataSourceTask(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("api_key").(string)
	client := NewClient(apiKey)
	return client, nil
}
