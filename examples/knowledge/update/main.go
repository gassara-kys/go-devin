package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	devin "github.com/gassara-kys/go-devin"
	"github.com/gassara-kys/go-devin/pkg/knowledge"
)

func main() {
	apiKey := os.Getenv("DEVIN_API_KEY")
	noteID := os.Getenv("DEVIN_KNOWLEDGE_ID")
	if apiKey == "" || noteID == "" {
		log.Fatal("DEVIN_API_KEY and DEVIN_KNOWLEDGE_ID are required")
	}

	client, err := devin.NewClient(apiKey)
	if err != nil {
		log.Fatalf("new client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	body := "Updated instructions for gathering weather."
	resp, err := client.Knowledge.Update(ctx, noteID, knowledge.UpdateRequest{
		Body: &body,
	})
	if err != nil {
		log.Fatalf("update knowledge: %v", err)
	}

	fmt.Printf("knowledge updated: %s (%s)\n", resp.ID, resp.Body)
}
