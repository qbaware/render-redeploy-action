package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

const (
	apiURL            = "https://api.render.com/v1"
	successStatusCode = 201
)

// Config represents the configuration for a Render client.
type Config struct {
	// APIKey is the API key used to authenticate with Render APIs.
	APIKey string
}

// Client represents a Render client.
type Client struct {
	Config Config
}

// NewClient creates a new Render client.
func NewClient(apiKey string) *Client {
	return &Client{
		Config: Config{
			APIKey: apiKey,
		},
	}
}

// Redeploy triggers a new deployment for a given service ID.
func (c *Client) Redeploy(serviceID string) error {
	url := fmt.Sprintf("%s/services/%s/deploys", apiURL, serviceID)
	body := "{}"

	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(body)))
	if err != nil {
		log.Fatalf("Failed to construct an HTTP request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Config.APIKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "call to Render APIs failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode != successStatusCode {
		return errors.Errorf("call to Render APIs failed, expected status code is %d but got %d", successStatusCode, resp.StatusCode)
	}

	return nil
}
