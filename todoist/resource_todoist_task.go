package todoist

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Task struct {
	ID        string    `json:"id"`
	ProjectID string    `json:"project_id"`
	Content   string `json:"content"`
	SectionID string `json:"section_id"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
	Labels []string `json:"labels"`
	Order int `json:"order"`
	Priority int `json:"priority"`
	Due struct {
		String string `json:"string"`
		Date string `json:"date"`
		IsRecurring bool `json:"is_recurring"`
		Datetime string `json:"datetime"`
		Timezone string `json:"timezone"`
	} `json:"due"`
	Url string `json:"url"`
	CommentCount int `json:"comment_count"`
	CreatedAt string `json:"created_at"`
	CreatorID string `json:"creator_id"`
	AssigneeID string `json:"assignee_id"`
	AssignerID string `json:"assigner_id"`
}

func ResourceTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskCreate,
		Read:   resourceTaskRead,
		Update: resourceTaskUpdate,
		Delete: resourceTaskDelete,
		Schema: map[string]*schema.Schema{
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

func resourceTaskCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	task := Task{
		Content:   d.Get("content").(string),
		ProjectID: d.Get("project_id").(string),
	}
	resp, err := client.resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(task).
		Post("/tasks")

	if err != nil {
		return err
	}

	var createdTask Task
	err = json.Unmarshal(resp.Body(), &createdTask)
	if err != nil {
		return err
	}

	d.SetId(createdTask.ID)
	d.Set("content", createdTask.Content)
	d.Set("project_id", createdTask.ProjectID)
	d.Set("section_id", createdTask.SectionID)
	d.Set("description", createdTask.Description)
	d.Set("is_completed", createdTask.IsCompleted)
	d.Set("labels", createdTask.Labels)
	d.Set("order", createdTask.Order)
	d.Set("priority", createdTask.Priority)
	d.Set("due", createdTask.Due)
	d.Set("url", createdTask.Url)
	d.Set("comment_count", createdTask.CommentCount)
	d.Set("created_at", createdTask.CreatedAt)
	d.Set("creator_id", createdTask.CreatorID)
	d.Set("assignee_id", createdTask.AssigneeID)
	d.Set("assigner_id", createdTask.AssignerID)

	return resourceTaskRead(d, m)
}

func resourceTaskRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	taskID := d.Id()
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

func resourceTaskUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	taskID := d.Id()
	task := Task{
		ID:        taskID,
		Content:   d.Get("content").(string),
		ProjectID: d.Get("project_id").(string),
	}
	_, err := client.resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(task).
		Post(fmt.Sprintf("/tasks/%s", taskID))

	if err != nil {
		return err
	}

	return resourceTaskRead(d, m)
}

func resourceTaskDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	taskID := d.Id()
	_, err := client.resty.R().
		Delete(fmt.Sprintf("/tasks/%s", taskID))

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}