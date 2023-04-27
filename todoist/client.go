package todoist

import (
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
