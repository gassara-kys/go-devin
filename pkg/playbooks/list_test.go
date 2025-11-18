package playbooks

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

func TestList(t *testing.T) {
	tests := []struct {
		name      string
		payload   []types.Playbook
		want      *ListResponse
		responder func([]types.Playbook) func(*http.Request) *http.Response
	}{
		{
			name:    "single playbook",
			payload: []types.Playbook{{PlaybookID: "pb-1", Title: "Example"}},
			want: &ListResponse{
				Playbooks: []types.Playbook{{PlaybookID: "pb-1", Title: "Example"}},
			},
			responder: func(data []types.Playbook) func(*http.Request) *http.Response {
				payload, _ := json.Marshal(data)
				return func(*http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(string(payload))),
						Header:     make(http.Header),
					}
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			svc := NewService(testutil.NewExecutor(t, tt.responder(tt.payload)), func(any) error { return nil })
			got, err := svc.List(context.Background())
			if err != nil {
				t.Fatalf("List error: %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
