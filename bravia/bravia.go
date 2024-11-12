package bravia

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// TODO: Create common response struct with error and result as they all have the same format
// TODO: Handle errors returned from the API

const (
	// headerAuthPSK is the header used for the pre-shared key authentication
	headerAuthPSK = "X-Auth-PSK"
)

type service struct {
	client *Client
}

// A Client is a client for interacting with the Sony Bravia API
type Client struct {
	client *http.Client // The HTTP client to use for requests

	BaseURL *url.URL // The base URL for the API

	// Services used for interacting with different parts of the API
	System *SystemService
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

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if v != nil {
		err := json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return resp, err
		}
	}

	return resp, nil
}

// copy returns a copy of the client
func (c *Client) copy() *Client {
	clone := Client{
		client:  &http.Client{},
		BaseURL: c.BaseURL,
	}

	// TODO: Preserve transport

	// TODO: Re-use services instead of creating new instances
	clone.System = &SystemService{client: &clone}

	return &clone
}

// authPSKTransport is an http.RoundTripper that adds the pre-shared key header to requests
type authPSKTransport struct {
	client  *http.Client
	BaseURL *url.URL
	PSK     string
}

// RoundTrip implements http.RoundTripper
func (t *authPSKTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(headerAuthPSK, t.PSK)
	return http.DefaultTransport.RoundTrip(req)
}

// WithAuthPSK returns a new client configured to use the given PSK for authentication
func (c *Client) WithAuthPSK(psk string) *Client {
	clone := c.copy()
	clone.client.Transport = &authPSKTransport{
		client:  clone.client,
		BaseURL: c.BaseURL,
		PSK:     psk,
	}
	return clone
}
