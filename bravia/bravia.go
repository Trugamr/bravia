package bravia

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// A Client is a client for interacting with the Sony Bravia API
type Client struct {
	client *http.Client // The HTTP client to use for requests

	BaseURL *url.URL // The base URL for the API

	// Services used for interacting with different parts of the API
	System *SystemService
}

type service struct {
	client *Client
}

func NewClient(baseURL *url.URL) *Client {
	client := &http.Client{}
	c := &Client{
		client:  client,
		BaseURL: baseURL,
	}

	// Initialize services
	c.System = &SystemService{client: c}

	return c
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if body != nil {
		err := json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), &buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
