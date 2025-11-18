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

	if err := client.Knowledge.Delete(ctx, noteID); err != nil {
		log.Fatalf("delete knowledge: %v", err)
	}
	log.Println("knowledge deleted")
}
