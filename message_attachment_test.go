package inboxroad_test

import (
	"fmt" //nolint:goimports,gofumpt,gci
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:goimports,gofumpt,gci
)

func TestMessageAttachment_SetName(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachment("", "", "")
	assert.Implements(t, (*inboxroad.MessageAttachmentInterface)(nil), h.SetName("file.txt"))
	assert.IsType(t, &inboxroad.MessageAttachment{}, h.SetName("file.txt"))
}

func TestMessageAttachment_GetName(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachment("", "", "")
	assert.Equal(t, "file.txt", h.SetName("file.txt").GetName())
}

func TestMessageAttachment_SetContent(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachment("", "", "")
	assert.Implements(t, (*inboxroad.MessageAttachmentInterface)(nil), h.SetContent("Test"))
	assert.IsType(t, &inboxroad.MessageAttachment{}, h.SetContent("Test"))
}

func TestMessageAttachment_GetContent(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachment("", "", "")
	assert.Equal(t, "Test", h.SetContent("Test").GetContent())
}

func TestMessageAttachment_GetContentAsBase64(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachment("", "", "")
	assert.Equal(t, "VGVzdA==", h.SetContent("Test").GetContentAsBase64())
}

func TestMessageAttachment_SetMimeType(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachment("", "", "")
	assert.Implements(t, (*inboxroad.MessageAttachmentInterface)(nil), h.SetMimeType("test/plain"))
	assert.IsType(t, &inboxroad.MessageAttachment{}, h.SetMimeType("test/plain"))
}

func TestMessageAttachment_GetMimeType(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachment("", "", "")
	assert.Equal(t, "text/plain", h.SetMimeType("text/plain").GetMimeType())
}

func TestMessageAttachment_ToArray(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachment("file.txt", "Test", "text/plain")
	data := h.ToArray()

	type strData struct {
		Key, Value string
	}

	stringValues := []strData{
		{"name", "file.txt"},
		{"content", "Test"},
		{"mimeType", "text/plain"},
	}

	for _, strVal := range stringValues {
		if val, ok := data[strVal.Key]; ok {
			assert.Equal(t, strVal.Value, val)
		} else {
			assert.Error(t, fmt.Errorf("%w: %s", errMissingMapKey, strVal.Key))
		}
	}
}

func TestMessageAttachment_ToInboxroadArray(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachment("file.txt", "Test", "text/plain")
	data := h.ToInboxroadArray()

	type strData struct {
		Key, Value string
	}

	stringValues := []strData{
		{"file_name", "file.txt"},
		{"file_data", "VGVzdA=="},
		{"mime_type", "text/plain"},
	}

	for _, strVal := range stringValues {
		if val, ok := data[strVal.Key]; ok {
			assert.Equal(t, strVal.Value, val)
		} else {
			assert.Error(t, fmt.Errorf("%w: %s", errMissingMapKey, strVal.Key))
		}
	}
}

func TestNewMessageAttachment(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachment("", "", "")
	assert.Implements(t, (*inboxroad.MessageAttachmentInterface)(nil), h)
	assert.IsType(t, &inboxroad.MessageAttachment{}, h)
}
