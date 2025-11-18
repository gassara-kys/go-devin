package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	devin "github.com/gassara-kys/go-devin"
	"github.com/gassara-kys/go-devin/pkg/sessions"
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

	resp, err := client.Sessions.List(ctx, &sessions.ListSessionsRequest{Limit: 5})
	if err != nil {
		log.Fatalf("list sessions: %v", err)
	}

	for _, s := range resp.Sessions {
		fmt.Printf("[%s] %s (%s)\n", s.Status, s.SessionID, s.Title)
	}
}
