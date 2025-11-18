package types

import (
	sessionsPkg "github.com/gassara-kys/go-devin/pkg/sessions"
)

type (
	ListSessionsRequest        = sessionsPkg.ListSessionsRequest
	ListSessionsResponse       = sessionsPkg.ListSessionsResponse
	CreateSessionRequest       = sessionsPkg.CreateSessionRequest
	CreateSessionResponse      = sessionsPkg.CreateSessionResponse
	SendMessageRequest         = sessionsPkg.SendMessageRequest
	SendMessageResponse        = sessionsPkg.SendMessageResponse
	UpdateTagsRequest          = sessionsPkg.UpdateTagsRequest
	UpdateTagsResponse         = sessionsPkg.UpdateTagsResponse
	UploadAttachmentRequest    = sessionsPkg.UploadAttachmentRequest
	UploadAttachmentResponse   = sessionsPkg.UploadAttachmentResponse
	DownloadAttachmentRequest  = sessionsPkg.DownloadAttachmentRequest
	DownloadAttachmentResponse = sessionsPkg.DownloadAttachmentResponse
)
