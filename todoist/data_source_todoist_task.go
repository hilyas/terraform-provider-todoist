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
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"section_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:    schema.TypeString,
				Optional: true,
			},
			"is_completed": {
				Type:    schema.TypeBool,
				Optional: true,
			},
			"labels": {
				Type:    schema.TypeList,
				Elem:    &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"order": {
				Type:    schema.TypeInt,
				Computed: true,
			},
			"priority": {
				Type:    schema.TypeInt,
				Optional: true,
			},
			"due": {
				Type:    schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:   &schema.Resource{
					Schema: map[string]*schema.Schema{
						"string": {
							Type: schema.TypeString,
							Required: true,
						},
						"date": {
							Type: schema.TypeString,
							Required: true,
						},
						"is_recurring": {
							Type: schema.TypeBool,
							Required: true,
						},
						"datetime": {
							Type: schema.TypeString,
							Optional: true,
						},
						"timezone": {
							Type: schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"url": {
				Type:    schema.TypeString,
				Computed: true,
			},
			"comment_count": {
				Type:    schema.TypeInt,
				Computed: true,
			},
			"created_at": {
				Type:    schema.TypeString,
				Computed: true,
			},
			"creator_id": {
				Type:    schema.TypeString,
				Computed: true,
			},
			"assignee_id": {
				Type:    schema.TypeString,
				Computed: true,
			},
			"assigner_id": {
				Type:    schema.TypeString,
				Computed: true,
			},
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

	var task Task
	err = json.Unmarshal(resp.Body(), &task)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s", task.ID))
	d.Set("content", task.Content)
	d.Set("project_id", task.ProjectID)
	d.Set("section_id", task.SectionID)
	d.Set("description", task.Description)
	d.Set("is_completed", task.IsCompleted)
	d.Set("labels", task.Labels)
	d.Set("order", task.Order)
	d.Set("priority", task.Priority)
	d.Set("due", task.Due)
	d.Set("url", task.Url)
	d.Set("comment_count", task.CommentCount)
	d.Set("created_at", task.CreatedAt)
	d.Set("creator_id", task.CreatorID)
	d.Set("assignee_id", task.AssigneeID)
	d.Set("assigner_id", task.AssignerID)

	return nil
}