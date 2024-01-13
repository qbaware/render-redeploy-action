package main

import (
	"log"
	"os"

	"github.com/qbaware/render-redeploy-action/internal/render"
)

func main() {
	renderAPIKey := os.Getenv("RENDER_API_KEY")
	renderServiceID := os.Getenv("RENDER_SERVICE_ID")

	if renderAPIKey == "" || renderServiceID == "" {
		log.Fatal("RENDER_API_KEY and RENDER_SERVICE_ID environment variables must be set")
	}

	log.Printf("Creating a new Render client")
	renderClient := render.NewClient(renderAPIKey)

	log.Printf("Triggering a redeploy for Render service with ID '%s'", renderServiceID)
	err := renderClient.Redeploy(renderServiceID)
	if err != nil {
		log.Fatalf("Failed to redeploy Render service with ID '%s': %v", renderServiceID, err)
	}

	log.Printf("Render service with ID '%s' has been successfully triggered to redeploy ✅", renderServiceID)
}
