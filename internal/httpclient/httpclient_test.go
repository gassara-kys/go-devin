package httpclient

import (
	"net/url"
	"testing"
)

func TestResolveURL(t *testing.T) {
	tests := []struct {
		name  string
		base  string
		path  string
		query url.Values
		want  string
	}{
		{
			name: "preserves versioned base",
			base: "https://api.devin.ai/v1",
			path: "/sessions",
			want: "https://api.devin.ai/v1/sessions",
		},
		{
			name:  "handles nested endpoints",
			base:  "https://api.devin.ai/v1",
			path:  "/sessions/devin-1/messages",
			query: url.Values{"limit": {"5"}},
			want:  "https://api.devin.ai/v1/sessions/devin-1/messages?limit=5",
		},
		{
			name:  "absolute path bypasses base",
			base:  "https://api.devin.ai/v1",
			path:  "https://uploads.devin.ai/files/foo",
			query: url.Values{"download": {"true"}},
			want:  "https://uploads.devin.ai/files/foo?download=true",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			exec := &Executor{cfg: &Config{BaseURL: tt.base}}
			got, err := exec.resolveURL(tt.path, tt.query)
			if err != nil {
				t.Fatalf("resolveURL error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("resolveURL mismatch: want %s got %s", tt.want, got)
			}
		})
	}
}
