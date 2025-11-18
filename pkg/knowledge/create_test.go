package knowledge

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/gassara-kys/go-devin/internal/testutil"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name      string
		request   CreateRequest
		responder func(*http.Request) *http.Response
		wantID    string
	}{
		{
			name: "basic",
			request: CreateRequest{
				Name: "Example",
				Body: "content",
			},
			wantID: "note-1",
			responder: func(r *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"id":"note-1","name":"Example","body":"content"}`)),
					Header:     make(http.Header),
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			svc := NewService(testutil.NewExecutor(t, tt.responder), func(any) error { return nil })
			resp, err := svc.Create(context.Background(), tt.request)
			if err != nil {
				t.Fatalf("Create returned error: %v", err)
			}
			if diff := cmp.Diff(tt.wantID, resp.ID); diff != "" {
				t.Fatalf("unexpected response (-want +got):\n%s", diff)
			}
		})
	}
}
