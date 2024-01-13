package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("RENDER_API_KEY")
	serviceID := os.Getenv("RENDER_SERVICE_ID")

	if apiKey == "" || serviceID == "" {
		log.Fatal("RENDER_API_KEY and RENDER_SERVICE_ID environment variables must be set")
	}

	url := fmt.Sprintf("https://api.render.com/v1/services/%s/deploys", serviceID)
	body := "{}"

	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(body)))
	if err != nil {
		log.Fatalf("Failed to construct an HTTP request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Call to Render APIs failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		log.Fatalf("Call to Render APIs failed, expected status code 201, got %d", resp.StatusCode)
	}

	log.Printf("Render service with ID '%s' has been successfully triggered to redeploy ✅", serviceID)
}
