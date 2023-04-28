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
	ParentID string `json:"parent_id"`
	Color string `json:"color"`
	IsFavorite bool `json:"is_favorite"`
	ViewStyle string `json:"view_style"`
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
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"color": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_favorite": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"view_style": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	project := Project{
		Name: d.Get("name").(string),
		ParentID: d.Get("parent_id").(string),
		Color: d.Get("color").(string),
		IsFavorite: d.Get("is_favorite").(bool),
		ViewStyle: d.Get("view_style").(string),
	}
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
	d.Set("name", createdProject.Name)
	d.Set("parent_id", createdProject.ParentID)
	d.Set("color", createdProject.Color)
	d.Set("is_favorite", createdProject.IsFavorite)
	d.Set("view_style", createdProject.ViewStyle)
	
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
	d.Set("parent_id", project.ParentID)
	d.Set("color", project.Color)
	d.Set("is_favorite", project.IsFavorite)
	d.Set("view_style", project.ViewStyle)

	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	project := Project{
		Name: d.Get("name").(string),
		ParentID: d.Get("parent_id").(string),
		Color: d.Get("color").(string),
		IsFavorite: d.Get("is_favorite").(bool),
		ViewStyle: d.Get("view_style").(string),
	}
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