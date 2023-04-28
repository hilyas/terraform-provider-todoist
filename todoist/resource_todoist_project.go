package todoist

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Project struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ParentID   string `json:"parent_id"`
	Color      string `json:"color"`
	IsFavorite bool   `json:"is_favorite"`
	ViewStyle  string `json:"view_style"`
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

	name := d.Get("name").(string)
	parentID := d.Get("parent_id").(string)
	color := d.Get("color").(string)
	isFavorite := d.Get("is_favorite").(bool)
	viewStyle := d.Get("view_style").(string)

	project, err := client.CreateProject(name, parentID, color, isFavorite, viewStyle)
	if err != nil {
		return err
	}

	d.SetId(project.ID)
	return resourceProjectRead(d, m)
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	project, err := client.GetProject(d.Id())
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

	name := d.Get("name").(string)
	parentID := d.Get("parent_id").(string)
	color := d.Get("color").(string)
	isFavorite := d.Get("is_favorite").(bool)
	viewStyle := d.Get("view_style").(string)

	_, err := client.UpdateProject(d.Id(), name, parentID, color, isFavorite, viewStyle)
	if err != nil {
		return err
	}

	return resourceProjectRead(d, m)
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteProject(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
