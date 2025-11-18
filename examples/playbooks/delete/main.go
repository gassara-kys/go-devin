package main

import (
	"context"
	"log"
	"os"
	"time"

	devin "github.com/gassara-kys/go-devin"
)

func main() {
	apiKey := os.Getenv("DEVIN_API_KEY")
	playbookID := os.Getenv("DEVIN_PLAYBOOK_ID")
	if apiKey == "" || playbookID == "" {
		log.Fatal("DEVIN_API_KEY and DEVIN_PLAYBOOK_ID are required")
	}

	client, err := devin.NewClient(apiKey)
	if err != nil {
		log.Fatalf("new client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.Playbooks.Delete(ctx, playbookID)
	if err != nil {
		log.Fatalf("delete playbook: %v", err)
	}

	log.Println(resp.Status)
}
