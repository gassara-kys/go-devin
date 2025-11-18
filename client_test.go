package devin

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/gassara-kys/go-devin/internal/testtransport"
	"github.com/gassara-kys/go-devin/pkg/sessions"
	"github.com/gassara-kys/go-devin/pkg/types"
)

func TestClientSessionsList(t *testing.T) {
	tests := []struct {
		name      string
		responder testtransport.RoundTripFunc
		request   *sessions.ListSessionsRequest
		want      *sessions.ListSessionsResponse
		wantErr   bool
	}{
		{
			name: "success",
			responder: func(*http.Request) *http.Response {
				payload := map[string]any{
					"sessions": []map[string]any{
						{
							"session_id": "devin-1",
							"status":     "running",
							"title":      "Example",
							"created_at": "2024-01-01T00:00:00Z",
							"updated_at": "2024-01-01T00:00:00Z",
						},
					},
				}
				body, _ := json.Marshal(payload)
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(string(body))),
					Header:     make(http.Header),
				}
			},
			request: &sessions.ListSessionsRequest{Limit: 5},
			want: &sessions.ListSessionsResponse{
				Sessions: []types.Session{{
					SessionID: "devin-1",
					Status:    "running",
					Title:     "Example",
					CreatedAt: "2024-01-01T00:00:00Z",
					UpdatedAt: "2024-01-01T00:00:00Z",
				}},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient("token", WithHTTPClient(testtransport.NewHTTPClient(tt.responder)))
			if err != nil {
				t.Fatalf("NewClient error: %v", err)
			}

			got, err := c.Sessions.List(context.Background(), tt.request)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("List error: %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("response mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
