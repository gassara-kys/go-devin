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
	if apiKey == "" {
		log.Fatal("DEVIN_API_KEY is required")
	}

	client, err := devin.NewClient(apiKey)
	if err != nil {
		log.Fatalf("new client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.Knowledge.Create(ctx, knowledge.CreateRequest{
		Name: "Weather instructions",
		Body: "Go to weather.com and fetch the temperature.",
	})
	if err != nil {
		log.Fatalf("create knowledge: %v", err)
	}

	fmt.Printf("knowledge created: %s (%s)\n", resp.ID, resp.Name)
}
