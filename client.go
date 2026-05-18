package go_boomi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

type Client struct {
	baseUrl     string
	username    string
	password    string
	rateLimiter *rate.Limiter
	context     context.Context
}

func NewClient(context context.Context, baseUrl, username, password string) *Client {
	return &Client{
		baseUrl,
		username,
		password,
		rate.NewLimiter(rate.Every(1*time.Second), 10),
		context,
	}
}

func (c *Client) Get(path string) ([]byte, error) {
	return c.Send("GET", path, nil, false)
}

func (c *Client) Post(path string, payload any) ([]byte, error) {
	return c.Send("POST", path, payload, false)
}

func (c *Client) Send(method string, path string, payload any, raw bool) ([]byte, error) {
	err := c.rateLimiter.Wait(c.context) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}

	var body io.Reader
	if payload != nil && raw == false {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
		body = bytes.NewBuffer(jsonData)
	} else if payload != nil && raw == true {
		strPayload, ok := payload.(string)
		if !ok {
			return nil, fmt.Errorf("payload must be a string when raw is true")
		}
		body = bytes.NewBuffer([]byte(strPayload))
	}

	fullURL, err := url.JoinPath(c.baseUrl, path)
	if err != nil {
		return nil, fmt.Errorf("failed to construct URL: %w", err)
	}

	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.username != "" && c.password != "" {
		req.SetBasicAuth(c.username, c.password)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}
