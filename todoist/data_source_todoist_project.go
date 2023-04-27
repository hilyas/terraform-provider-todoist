package todoist

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProjectRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceProjectRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	projectID := d.Get("project_id").(string)
	resp, err := client.resty.R().Get(fmt.Sprintf("/projects/%s", projectID))
	if err != nil {
		return err
	}

	var project Project
	err = json.Unmarshal(resp.Body(), &project)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d", project.ID))
	d.Set("name", project.Name)

	return nil
}
