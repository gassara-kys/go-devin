# go-devin

Go SDK for the [Devin API v1](https://docs.devin.ai/api-reference/overview). It exposes first-class packages for each domain (sessions, secrets, knowledge, playbooks) and a root `devin` package for client creation.

## Features

- Ergonomic `devin.NewClient` with pluggable HTTP client, logger, user agent, and retry configuration.
- Strongly typed request/response structs per endpoint under `pkg/<domain>`, aligned with the official docs.
- Shared HTTP executor (`internal/httpclient`) with automatic retries (408/425/429/5xx), request logging, and bearer authentication.
- Gin-compatible validation and `go-querystring`-based query encoding.
- Complete runnable samples under `examples/<domain>/<endpoint>` that consume `DEVIN_API_KEY`.

## Installation

```bash
go get github.com/gassara-kys/go-devin
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "os"

    devin "github.com/gassara-kys/go-devin"
    "github.com/gassara-kys/go-devin/pkg/sessions"
)

func main() {
    apiKey := os.Getenv("DEVIN_API_KEY")
    if apiKey == "" {
        panic("DEVIN_API_KEY is not set")
    }

    client, err := devin.NewClient(apiKey)
    if err != nil {
        panic(err)
    }

    ctx := context.Background()
    summaries, err := client.Sessions.List(ctx, &sessions.ListSessionsRequest{Limit: 5})
    if err != nil {
        panic(err)
    }

    for _, sess := range summaries.Sessions {
        fmt.Println(sess.SessionID)
    }
}
```

## Session Details

`SessionDetail.Messages` is populated by `Sessions.Get` and includes the transcript metadata (`event_id`, `message`, `timestamp`, `type`, `origin`, `user_id`, `username`).

```go
detail, err := client.Sessions.Get(ctx, "devin-123")
if err != nil {
    panic(err)
}
for _, msg := range detail.Messages {
    fmt.Printf("%s %s\n", msg.Timestamp, msg.Message)
}
```

## Examples

Each endpoint has an executable sample. For instance:

```bash
DEVIN_API_KEY=xxx go run ./examples/sessions/list
DEVIN_API_KEY=xxx DEVIN_SESSION_ID=devin-123 go run ./examples/sessions/get
DEVIN_API_KEY=xxx DEVIN_SESSION_ID=devin-123 go run ./examples/sessions/send_message
```

## Development

```bash
make lint   # golangci-lint run ./...
make test   # go test ./...
make build  # go build ./...
```

## License

MIT
