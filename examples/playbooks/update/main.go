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

	req := playbooks.UpdateRequest{Title: "Updated Playbook Title"}
	resp, err := client.Playbooks.Update(ctx, playbookID, req)
	if err != nil {
		log.Fatalf("update playbook: %v", err)
	}

	fmt.Println(resp.Status)
}
