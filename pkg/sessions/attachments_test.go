package sessions

import (
	"bytes"
	"context"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/gassara-kys/go-devin/internal/testutil"
)

func TestUploadAttachment(t *testing.T) {
	var captured string
	tests := []struct {
		name      string
		request   UploadAttachmentRequest
		responder func(*http.Request) *http.Response
		want      *UploadAttachmentResponse
		wantBody  string
		expectErr bool
	}{
		{
			name: "single file",
			request: UploadAttachmentRequest{
				FileName: "notes.txt",
				Reader:   strings.NewReader("hello"),
			},
			want:     &UploadAttachmentResponse{AttachmentID: "file-123"},
			wantBody: "hello",
			responder: func(r *http.Request) *http.Response {
				ct := r.Header.Get("Content-Type")
				_, params, err := mime.ParseMediaType(ct)
				if err != nil {
					t.Fatalf("ParseMediaType: %v", err)
				}
				reader := multipart.NewReader(r.Body, params["boundary"])
				part, err := reader.NextPart()
				if err != nil {
					t.Fatalf("NextPart: %v", err)
				}
				data, _ := io.ReadAll(part)
				captured = string(data)
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("file-123")),
					Header:     make(http.Header),
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			svc := NewService(testutil.NewExecutor(t, tt.responder), func(any) error { return nil })
			got, err := svc.UploadAttachment(context.Background(), tt.request)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("UploadAttachment error: %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("response mismatch (-want +got):\n%s", diff)
			}
			if captured != tt.wantBody {
				t.Fatalf("unexpected payload: want %q got %q", tt.wantBody, captured)
			}
			captured = ""
		})
	}
}

func TestDownloadAttachment(t *testing.T) {
	tests := []struct {
		name      string
		request   DownloadAttachmentRequest
		responder func(*http.Request) *http.Response
		want      *DownloadAttachmentResponse
	}{
		{
			name: "basic",
			request: DownloadAttachmentRequest{
				AttachmentID: "att-1",
				FileName:     "report.txt",
			},
			responder: func(r *http.Request) *http.Response {
				if r.URL.Path != "/attachments/att-1/report.txt" {
					t.Fatalf("unexpected path %s", r.URL.Path)
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString("payload")),
					Header:     make(http.Header),
				}
			},
			want: &DownloadAttachmentResponse{
				AttachmentID: "att-1",
				FileName:     "report.txt",
				Content:      []byte("payload"),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			svc := NewService(testutil.NewExecutor(t, tt.responder), func(any) error { return nil })
			got, err := svc.DownloadAttachment(context.Background(), tt.request)
			if err != nil {
				t.Fatalf("DownloadAttachment error: %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("unexpected response (-want +got):\n%s", diff)
			}
		})
	}
}
