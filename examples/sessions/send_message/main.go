package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	devin "github.com/gassara-kys/go-devin"
	"github.com/gassara-kys/go-devin/pkg/sessions"
)

func main() {
	apiKey := os.Getenv("DEVIN_API_KEY")
	sessionID := os.Getenv("DEVIN_SESSION_ID")
	if apiKey == "" || sessionID == "" {
		log.Fatal("DEVIN_API_KEY and DEVIN_SESSION_ID are required")
	}

	client, err := devin.NewClient(apiKey)
	if err != nil {
		log.Fatalf("new client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := sessions.SendMessageRequest{
		Message: "Please describe your current progress.",
	}
	resp, err := client.Sessions.SendMessage(ctx, sessionID, req)
	if err != nil {
		log.Fatalf("send message: %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("encode json: %v", err)
	}
}
