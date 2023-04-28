package todoist

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	resty *resty.Client
}

func NewClient(apiKey string) *Client {
	client := resty.New()

	client.SetHostURL("https://api.todoist.com/rest/v2").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+apiKey)

	return &Client{
		resty: client,
	}
}

func (c *Client) CreateProject(name, parentID, color string, isFavorite bool, viewStyle string) (*Project, error) {
	projectData := &Project{
		Name:       name,
		ParentID:   parentID,
		Color:      color,
		IsFavorite: isFavorite,
		ViewStyle:  viewStyle,
	}

	resp, err := c.resty.R().SetBody(projectData).Post("/projects")
	if err != nil {
		return nil, err
	}

	var project Project
	err = json.Unmarshal(resp.Body(), &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) GetProject(projectID string) (*Project, error) {
	resp, err := c.resty.R().Get("/projects/" + projectID)
	if err != nil {
		return nil, err
	}

	var project Project
	err = json.Unmarshal(resp.Body(), &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) UpdateProject(projectID, name, parentID, color string, isFavorite bool, viewStyle string) (*Project, error) {
	projectData := &Project{
		Name:       name,
		ParentID:   parentID,
		Color:      color,
		IsFavorite: isFavorite,
		ViewStyle:  viewStyle,
	}

	resp, err := c.resty.R().SetBody(projectData).Post("/projects/" + projectID)
	if err != nil {
		return nil, err
	}

	var project Project
	err = json.Unmarshal(resp.Body(), &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) DeleteProject(projectID string) error {
	_, err := c.resty.R().Delete("/projects/" + projectID)
	return err
}

func (c *Client) CreateTask(content, projectID, description string, isCompleted bool, labels []string, priority int, dueString string) (*Task, error) {
	taskData := &Task{
		Content:     content,
		ProjectID:   projectID,
		Description: description,
		IsCompleted: isCompleted,
		Labels:      labels,
		Priority:    priority,
		DueString:   dueString,
	}

	resp, err := c.resty.R().SetBody(taskData).Post("/tasks")
	if err != nil {
		return nil, err
	}

	var task Task
	err = json.Unmarshal(resp.Body(), &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (c *Client) GetTask(taskID string) (*Task, error) {
	resp, err := c.resty.R().Get("/tasks/" + taskID)
	if err != nil {
		return nil, err
	}

	var task Task
	err = json.Unmarshal(resp.Body(), &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (c *Client) UpdateTask(taskID, content, projectID, description string, isCompleted bool, labels []string, priority int, dueString string) (*Task, error) {
	taskData := &Task{
		Content:     content,
		ProjectID:   projectID,
		Description: description,
		IsCompleted: isCompleted,
		Labels:      labels,
		Priority:    priority,
		DueString:   dueString,
	}

	resp, err := c.resty.R().SetBody(taskData).Post("/tasks/" + taskID)
	if err != nil {
		return nil, err
	}

	var task Task
	err = json.Unmarshal(resp.Body(), &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (c *Client) DeleteTask(taskID string) error {
	_, err := c.resty.R().Delete("/tasks/" + taskID)
	return err
}
