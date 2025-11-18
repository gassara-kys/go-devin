package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	devin "github.com/gassara-kys/go-devin"
	"github.com/gassara-kys/go-devin/pkg/playbooks"
)

func main() {
	apiKey := os.Getenv("DEVIN_API_KEY")
	if apiKey == "" {
		log.Fatal("DEVIN_API_KEY is required")
	}

	client, err := devin.NewClient(apiKey)
	if err != nil {
		log.Fatalf("new client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.Playbooks.Create(ctx, playbooks.CreateRequest{
		Title: "Notify on deployment",
		Body:  "When a deployment completes, send a Slack message.",
	})
	if err != nil {
		log.Fatalf("create playbook: %v", err)
	}

	fmt.Printf("playbook created: %s\n", resp.PlaybookID)
}
