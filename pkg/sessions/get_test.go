package sessions

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/gassara-kys/go-devin/internal/testutil"
	"github.com/gassara-kys/go-devin/pkg/types"
)

func TestServiceGet(t *testing.T) {
	origin := "user"
	userID := "user-1"
	username := "alice"

	tests := []struct {
		name      string
		sessionID string
		responder func(*http.Request) *http.Response
		want      *types.SessionDetail
		wantErr   bool
	}{
		{
			name:      "basic",
			sessionID: "devin-1",
			responder: func(r *http.Request) *http.Response {
				if r.URL.Path != "/sessions/devin-1" {
					t.Fatalf("unexpected path %s", r.URL.Path)
				}
				payload := map[string]any{
					"session_id": "devin-1",
					"status":     "working",
					"title":      "Test session",
					"created_at": "2024-01-01T00:00:00Z",
					"updated_at": "2024-01-02T00:00:00Z",
					"messages": []map[string]any{
						{
							"event_id":  "evt-1",
							"message":   "hello",
							"timestamp": "2024-01-01T00:01:00Z",
							"type":      "text",
							"origin":    "user",
							"user_id":   "user-1",
							"username":  "alice",
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
			want: &types.SessionDetail{
				Session: types.Session{
					SessionID: "devin-1",
					Status:    "working",
					Title:     "Test session",
					CreatedAt: "2024-01-01T00:00:00Z",
					UpdatedAt: "2024-01-02T00:00:00Z",
				},
				Messages: []types.SessionMessage{
					{
						EventID:   "evt-1",
						Message:   "hello",
						Timestamp: "2024-01-01T00:01:00Z",
						Type:      "text",
						Origin:    &origin,
						UserID:    &userID,
						Username:  &username,
					},
				},
			},
		},
		{
			name:      "missing session id",
			sessionID: "",
			responder: func(r *http.Request) *http.Response {
				t.Fatalf("unexpected request %s", r.URL.Path)
				return nil
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			svc := NewService(testutil.NewExecutor(t, tt.responder), func(any) error { return nil })
			got, err := svc.Get(context.Background(), tt.sessionID)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Get error: %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("response mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
