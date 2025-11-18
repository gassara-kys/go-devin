package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	devin "github.com/gassara-kys/go-devin"
	"github.com/gassara-kys/go-devin/pkg/secrets"
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

	req := secrets.CreateRequest{
		Type:  "api_key",
		Key:   "SERVICE_TOKEN",
		Value: "secret-value",
		Note:  "example secret",
	}
	resp, err := client.Secrets.Create(ctx, req)
	if err != nil {
		log.Fatalf("create secret: %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("encode json: %v", err)
	}
}
