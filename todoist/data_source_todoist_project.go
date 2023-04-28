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
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"color": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_favorite": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"view_style": {
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

	d.SetId(fmt.Sprintf("%s", project.ID))
	d.Set("name", project.Name)
	d.Set("parent_id", project.ParentID)
	d.Set("color", project.Color)
	d.Set("is_favorite", project.IsFavorite)
	d.Set("view_style", project.ViewStyle)

	return nil
}
