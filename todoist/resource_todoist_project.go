package todoist

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Add the Project struct
type Project struct {
	ID   string  `json:"id"`
	Name string `json:"name"`
	
}

func ResourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	project := Project{Name: d.Get("name").(string)}
	resp, err := client.resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(project).
		Post("/projects")

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Read response body: %s", resp.Body())

	var createdProject Project
	err = json.Unmarshal(resp.Body(), &createdProject)
	if err != nil {
		return err
	}

	d.SetId(createdProject.ID)

	return resourceProjectRead(d, m)
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	resp, err := client.resty.R().Get("/projects/" + d.Id())
	if err != nil {
		return err
	}

	var project Project
	err = json.Unmarshal(resp.Body(), &project)
	if err != nil {
		return err
	}

	d.Set("name", project.Name)

	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	project := Project{Name: d.Get("name").(string)}
	resp, err := client.resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(project).
		Post("/projects/" + d.Id())

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Update response body: %s", resp.Body())

	return resourceProjectRead(d, m)
}


func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	_, err := client.resty.R().Delete("/projects/" + d.Id())
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}