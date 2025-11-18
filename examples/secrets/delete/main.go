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
	secretID := os.Getenv("DEVIN_SECRET_ID")
	if apiKey == "" || secretID == "" {
		log.Fatal("DEVIN_API_KEY and DEVIN_SECRET_ID are required")
	}

	client, err := devin.NewClient(apiKey)
	if err != nil {
		log.Fatalf("new client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.Secrets.Delete(ctx, secretID); err != nil {
		log.Fatalf("delete secret: %v", err)
	}
	log.Println("secret deleted")
}
