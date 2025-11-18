package sessions

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type UploadAttachmentRequest struct {
	FileName string    `binding:"required"`
	Reader   io.Reader `binding:"required"`
}

type UploadAttachmentResponse struct {
	AttachmentID string `json:"attachment_id"`
}

func (s *Service) UploadAttachment(ctx context.Context, req UploadAttachmentRequest) (*UploadAttachmentResponse, error) {
	if err := s.validateStruct(&req); err != nil {
		return nil, err
	}
	if req.Reader == nil {
		return nil, errors.New("reader is required")
	}
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	part, err := writer.CreateFormFile("file", path.Base(req.FileName))
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, req.Reader); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	body := buf.Bytes()
	respBody, err := s.doBytes(ctx, http.MethodPost, "/files", nil, body, writer.FormDataContentType(), "text/plain")
	if err != nil {
		return nil, err
	}
	return &UploadAttachmentResponse{AttachmentID: strings.TrimSpace(string(respBody))}, nil
}

type DownloadAttachmentRequest struct {
	AttachmentID string `binding:"required"`
	FileName     string `binding:"required"`
}

type DownloadAttachmentResponse struct {
	AttachmentID string
	FileName     string
	Content      []byte
}

func (s *Service) DownloadAttachment(ctx context.Context, req DownloadAttachmentRequest) (*DownloadAttachmentResponse, error) {
	if err := s.validateStruct(&req); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/attachments/%s/%s", url.PathEscape(req.AttachmentID), url.PathEscape(req.FileName))
	content, err := s.doBytes(ctx, http.MethodGet, path, nil, nil, "", "")
	if err != nil {
		return nil, err
	}
	return &DownloadAttachmentResponse{
		AttachmentID: req.AttachmentID,
		FileName:     req.FileName,
		Content:      content,
	}, nil
}
