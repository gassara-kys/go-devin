package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	devin "github.com/gassara-kys/go-devin"
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

	resp, err := client.Playbooks.List(ctx)
	if err != nil {
		log.Fatalf("list playbooks: %v", err)
	}
	for _, pb := range resp.Playbooks {
		fmt.Printf("%s (%s)\n", pb.PlaybookID, pb.Title)
	}
}
