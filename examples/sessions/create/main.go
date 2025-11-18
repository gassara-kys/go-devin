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

	req := sessions.CreateSessionRequest{
		Prompt: "Summarize repository README",
		Tags:   []string{"example"},
	}
	resp, err := client.Sessions.Create(ctx, req)
	if err != nil {
		log.Fatalf("create session: %v", err)
	}

	fmt.Printf("session created: %s (status=%s)\n", resp.SessionID, resp.Status)
}
