package secrets

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
	var captured string
	tests := []struct {
		name      string
		request   CreateRequest
		responder func(*http.Request) *http.Response
		want      *CreateResponse
		wantBody  string
	}{
		{
			name: "basic secret",
			request: CreateRequest{
				Type:  "api_key",
				Key:   "OPENAI",
				Value: "secret",
			},
			want: &CreateResponse{ID: "sec-1"},
			responder: func(r *http.Request) *http.Response {
				data, _ := io.ReadAll(r.Body)
				captured = string(data)
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"id":"sec-1"}`)),
					Header:     make(http.Header),
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			svc := NewService(testutil.NewExecutor(t, tt.responder), func(any) error { return nil })
			got, err := svc.Create(context.Background(), tt.request)
			if err != nil {
				t.Fatalf("Create error: %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("response mismatch (-want +got):\n%s", diff)
			}
			if !strings.Contains(captured, `"key":"OPENAI"`) {
				t.Fatalf("payload missing key: %s", captured)
			}
			captured = ""
		})
	}
}
