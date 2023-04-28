package todoist

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Task struct {
	ID          string   `json:"id"`
	Content     string   `json:"content"`
	ProjectID   string   `json:"project_id"`
	// SectionID   string   `json:"section_id"`
	Description string   `json:"description"`
	IsCompleted bool     `json:"is_completed"`
	Labels      []string `json:"labels"`
	Priority    int      `json:"priority"`
	// DueString  string   `json:"due_string"`
	// DueDate    string   `json:"due_date"`
	// DueDatetime string   `json:"due_datetime"`
	// DueLang   string   `json:"due_lang"`
	// Due         *Due     `json:"due"`
}

// type Due struct {
// 	String      string `json:"string"`
// 	Date        string `json:"date"`
// 	IsRecurring bool   `json:"is_recurring"`
// 	Datetime    string `json:"datetime"`
// 	Timezone    string `json:"timezone"`
// }


func ResourceTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceTodoistTaskCreate,
		Read:   resourceTodoistTaskRead,
		Update: resourceTodoistTaskUpdate,
		Delete: resourceTodoistTaskDelete,

		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// "section_id": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_completed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// "due_string": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "due_date": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "due_datetime": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "due_lang": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "due": {
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	MaxItems: 1,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"string": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"date": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"is_recurring": {
			// 				Type:     schema.TypeBool,
			// 				Required: true,
			// 			},
			// 			"datetime": {
			// 				Type:     schema.TypeString,
			// 				Optional: true,
			// 			},
			// 			"timezone": {
			// 				Type:     schema.TypeString,
			// 				Optional: true,
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

func resourceTodoistTaskCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	task, err := client.CreateTask(
		d.Get("content").(string),
		d.Get("project_id").(string),
		// d.Get("section_id").(string),
		d.Get("description").(string),
		d.Get("is_completed").(bool),
		expandLabels(d.Get("labels")),
		d.Get("priority").(int),
		// expandDue(d.Get("due")),
	)
	if err != nil {
		return fmt.Errorf("Error creating Todoist task: %s", err)
	}

	if task.ID == "" {
        return fmt.Errorf("Error creating Todoist task: Task ID is empty")
    }
	d.SetId(task.ID)

	return resourceTodoistTaskRead(d, m)
}

func resourceTodoistTaskRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	task, err := client.GetTask(d.Id())
	if err != nil {
		return fmt.Errorf("Error reading Todoist task: %s", err)
	}

	d.Set("content", task.Content)
	d.Set("project_id", task.ProjectID)
	// d.Set("section_id", task.SectionID)
	d.Set("description", task.Description)
	d.Set("is_completed", task.IsCompleted)
	d.Set("labels", flattenLabels(task.Labels))
	d.Set("priority", task.Priority)
	// d.Set("due", flattenDue(task.Due))

	return nil
}

func resourceTodoistTaskUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	task, err := client.UpdateTask(
		d.Id(),
		d.Get("content").(string),
		d.Get("project_id").(string),
		// d.Get("section_id").(string),
		d.Get("description").(string),
		d.Get("is_completed").(bool),
		expandLabels(d.Get("labels")),
		d.Get("priority").(int),
		// expandDue(d.Get("due")),
	)
	if err != nil {
		return fmt.Errorf("Error updating Todoist task: %s", err)
	}

	if task.ID == "" {
        return fmt.Errorf("Error updating Todoist task: Task ID is empty")
    }
	d.SetId(task.ID)

	return resourceTodoistTaskRead(d, m)
}

func resourceTodoistTaskDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteTask(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Todoist task: %s", err)
	}

	return nil
}

func expandLabels(labels interface{}) []string {
	labelList := labels.([]interface{})
	expandedLabels := make([]string, len(labelList))

	for i, label := range labelList {
		expandedLabels[i] = label.(string)
	}

	return expandedLabels
}

// func expandDue(due interface{}) *Due {
// 	dueList := due.([]interface{})
// 	if len(dueList) == 0 {
// 		return nil
// 	}

// 	dueMap := dueList[0].(map[string]interface{})

// 	dueStruct := &Due{}

// 	if value, ok := dueMap["string"].(string); ok {
// 		dueStruct.String = value
// 	}
// 	if value, ok := dueMap["date"].(string); ok {
// 		dueStruct.Date = value
// 	}
// 	if value, ok := dueMap["is_recurring"].(bool); ok {
// 		dueStruct.IsRecurring = value
// 	}
// 	if value, ok := dueMap["datetime"].(string); ok {
// 		dueStruct.Datetime = value
// 	}
// 	if value, ok := dueMap["timezone"].(string); ok {
// 		dueStruct.Timezone = value
// 	}

// 	return dueStruct
// }

func flattenLabels(labels []string) []interface{} {
	flattened := make([]interface{}, len(labels))

	for i, label := range labels {
		flattened[i] = label
	}

	return flattened
}

// func flattenDue(due *Due) []interface{} {
// 	if due == nil {
// 		return []interface{}{}
// 	}

// 	return []interface{}{
// 		map[string]interface{}{
// 			"string":       due.String,
// 			"date":         due.Date,
// 			"is_recurring": due.IsRecurring,
// 			"datetime":     due.Datetime,
// 			"timezone":     due.Timezone,
// 		},
// 	}
// }

