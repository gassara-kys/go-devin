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

func TestServiceList(t *testing.T) {
	tests := []struct {
		name      string
		request   *ListSessionsRequest
		responder func(*http.Request) *http.Response
		want      *ListSessionsResponse
	}{
		{
			name:    "limit 5",
			request: &ListSessionsRequest{Limit: 5},
			responder: func(r *http.Request) *http.Response {
				if got := r.URL.Query().Get("limit"); got != "5" {
					t.Fatalf("expected limit=5 got=%s", got)
				}
				payload := map[string]any{
					"sessions": []map[string]any{
						{
							"session_id": "devin-1",
							"status":     "running",
							"title":      "Test",
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
			want: &ListSessionsResponse{
				Sessions: []types.Session{{
					SessionID: "devin-1",
					Status:    "running",
					Title:     "Test",
					CreatedAt: "2024-01-01T00:00:00Z",
					UpdatedAt: "2024-01-01T00:00:00Z",
				}},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			transport := testutil.NewExecutor(t, tt.responder)
			svc := NewService(transport, func(any) error { return nil })

			got, err := svc.List(context.Background(), tt.request)
			if err != nil {
				t.Fatalf("List returned error: %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("response mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
