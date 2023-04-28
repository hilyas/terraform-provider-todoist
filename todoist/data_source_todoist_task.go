package todoist

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskRead,
		Schema: map[string]*schema.Schema{
			"task_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"section_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_completed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"due_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"due_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"due_datetime": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"due_lang": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// "due": {
			// 	Type:     schema.TypeList,
			// 	Computed: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"string": {
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 			"date": {
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 			"is_recurring": {
			// 				Type:     schema.TypeBool,
			// 				Computed: true,
			// 			},
			// 			"datetime": {
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 			"timezone": {
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

func dataSourceTaskRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	taskID := d.Get("task_id").(string)
	resp, err := client.resty.R().Get(fmt.Sprintf("/tasks/%s", taskID))
	if err != nil {
		return err
	}

	fmt.Printf("JSON response: %s\n", string(resp.Body())) // print the JSON response

	var task Task
	err = json.Unmarshal(resp.Body(), &task)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s", task.ID))
	d.Set("content", task.Content)
	d.Set("project_id", task.ProjectID)
	// d.Set("section_id", task.SectionID)
	d.Set("description", task.Description)
	d.Set("is_completed", task.IsCompleted)
	d.Set("labels", task.Labels)
	d.Set("priority", task.Priority)

	// // Set due attribute
	// due := map[string]interface{}{
	// 	"string":       task.Due.String,
	// 	"date":         task.Due.Date,
	// 	"is_recurring": task.Due.IsRecurring,
	// 	"datetime":     task.Due.Datetime,
	// 	"timezone":     task.Due.Timezone,
	// }
	// if err := d.Set("due", []interface{}{due}); err != nil {
	// 	return err
	// }

	return nil
}
