package types

import (
	sessionsPkg "github.com/gassara-kys/go-devin/pkg/sessions"
)

type (
	// ListSessionsRequest documents the request payload for listing sessions.
	ListSessionsRequest = sessionsPkg.ListSessionsRequest
	// ListSessionsResponse documents the response payload for listing sessions.
	ListSessionsResponse = sessionsPkg.ListSessionsResponse
	// CreateSessionRequest documents the request payload for creating sessions.
	CreateSessionRequest = sessionsPkg.CreateSessionRequest
	// CreateSessionResponse documents the response payload for creating sessions.
	CreateSessionResponse = sessionsPkg.CreateSessionResponse
	// SendMessageRequest documents the request payload for sending session messages.
	SendMessageRequest = sessionsPkg.SendMessageRequest
	// SendMessageResponse documents the response payload for sending session messages.
	SendMessageResponse = sessionsPkg.SendMessageResponse
	// UpdateTagsRequest documents the request payload for updating session tags.
	UpdateTagsRequest = sessionsPkg.UpdateTagsRequest
	// UpdateTagsResponse documents the response payload for updating session tags.
	UpdateTagsResponse = sessionsPkg.UpdateTagsResponse
	// UploadAttachmentRequest documents the request payload for uploading attachments.
	UploadAttachmentRequest = sessionsPkg.UploadAttachmentRequest
	// UploadAttachmentResponse documents the response payload for uploading attachments.
	UploadAttachmentResponse = sessionsPkg.UploadAttachmentResponse
	// DownloadAttachmentRequest documents the request payload for downloading attachments.
	DownloadAttachmentRequest = sessionsPkg.DownloadAttachmentRequest
	// DownloadAttachmentResponse documents the response payload for downloading attachments.
	DownloadAttachmentResponse = sessionsPkg.DownloadAttachmentResponse
)
