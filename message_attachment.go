package inboxroad

import "encoding/base64"

type MessageAttachmentInterface interface {
	SetName(name string) MessageAttachmentInterface
	GetName() string
	SetContent(content string) MessageAttachmentInterface
	GetContent() string
	SetMimeType(mimeType string) MessageAttachmentInterface
	GetMimeType() string
	GetContentAsBase64() string
	ToMap() map[string]string
	ToInboxroadMap() map[string]string
}

type MessageAttachment struct {
	name, content, mimeType string
}

func (m *MessageAttachment) SetName(name string) MessageAttachmentInterface {
	m.name = name

	return m
}

func (m MessageAttachment) GetName() string {
	return m.name
}

func (m *MessageAttachment) SetContent(content string) MessageAttachmentInterface {
	m.content = content

	return m
}

func (m MessageAttachment) GetContent() string {
	return m.content
}

func (m MessageAttachment) GetContentAsBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(m.content))
}

func (m *MessageAttachment) SetMimeType(mimeType string) MessageAttachmentInterface {
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	m.mimeType = mimeType

	return m
}

func (m MessageAttachment) GetMimeType() string {
	return m.mimeType
}

func (m MessageAttachment) ToMap() map[string]string {
	return map[string]string{
		"name":     m.GetName(),
		"content":  m.GetContent(),
		"mimeType": m.GetMimeType(),
	}
}

func (m MessageAttachment) ToInboxroadMap() map[string]string {
	return map[string]string{
		"filename":  m.GetName(),
		"file_data": m.GetContentAsBase64(),
		"mime_type": m.GetMimeType(),
	}
}

func NewMessageAttachment(name, content, mimeType string) MessageAttachmentInterface {
	return (&MessageAttachment{}). //nolint:exhaustivestruct
					SetName(name).
					SetContent(content).
					SetMimeType(mimeType)
}
