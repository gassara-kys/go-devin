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

	req := sessions.UpdateTagsRequest{
		Tags: []string{"demo", "api"},
	}
	resp, err := client.Sessions.UpdateTags(ctx, sessionID, req)
	if err != nil {
		log.Fatalf("update tags: %v", err)
	}

	fmt.Println(resp.Detail)
}
