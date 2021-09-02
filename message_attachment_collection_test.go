package inboxroad_test

import ( //nolint:nolintlint,goimports,gofumpt,gci
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:nolintlint,goimports,gofumpt,gci
)

func TestMessageAttachmentCollection_Add(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachmentCollection().
		Add(inboxroad.NewMessageAttachment("file-1.txt", "Test 1", "text/plain")).
		Add(inboxroad.NewMessageAttachment("file-2.txt", "Test 2", "text/plain"))

	assert.Implements(t, (*inboxroad.MessageAttachmentCollectionInterface)(nil), h)
	assert.IsType(t, &inboxroad.MessageAttachmentCollection{}, h)
}

func TestMessageAttachmentCollection_GetItems(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachmentCollection().
		Add(inboxroad.NewMessageAttachment("file-1.txt", "Test 1", "text/plain")).
		Add(inboxroad.NewMessageAttachment("file-2.txt", "Test 2", "text/plain"))

	items := h.GetItems()
	assert.Equal(t, 2, len(items))
	assert.Equal(t, items[0].GetName(), "file-1.txt")
	assert.Equal(t, items[0].GetContent(), "Test 1")
	assert.Equal(t, items[0].GetMimeType(), "text/plain")
}

func TestMessageAttachmentCollection_ToSliceMap(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachmentCollection().
		Add(inboxroad.NewMessageAttachment("file-1.txt", "Test 1", "text/plain")).
		Add(inboxroad.NewMessageAttachment("file-2.txt", "Test 2", "text/plain"))

	data := h.ToSliceMap()

	result := inboxroad.SliceStringMap{
		{"name": "file-1.txt", "content": "Test 1", "mimeType": "text/plain"},
		{"name": "file-2.txt", "content": "Test 2", "mimeType": "text/plain"},
	}

	assert.Equal(t, result, data)
}

func TestMessageAttachmentCollection_ToInboxroadSliceMap(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachmentCollection().
		Add(inboxroad.NewMessageAttachment("file-1.txt", "Test 1", "text/plain")).
		Add(inboxroad.NewMessageAttachment("file-2.txt", "Test 2", "text/plain"))

	data := h.ToInboxroadSliceMap()

	result := inboxroad.SliceStringMap{
		{"filename": "file-1.txt", "file_data": "VGVzdCAx", "mime_type": "text/plain"},
		{"filename": "file-2.txt", "file_data": "VGVzdCAy", "mime_type": "text/plain"},
	}

	assert.Equal(t, result, data)
}

func TestNewMessageAttachmentCollection(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachmentCollection()
	assert.Implements(t, (*inboxroad.MessageAttachmentCollectionInterface)(nil), h)
	assert.IsType(t, &inboxroad.MessageAttachmentCollection{}, h)
}

func TestNewMessageAttachmentCollectionFromSliceMap(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageAttachmentCollectionFromSliceMap(inboxroad.SliceStringMap{
		{"name": "file-1.txt", "content": "Test 1", "mimeType": "text/plain"},
		{"name": "file-2.txt", "content": "Test 2", "mimeType": "text/plain"},
	})

	assert.Equal(t, 2, len(h.GetItems()))
	assert.Equal(t, h.GetItems()[0].GetName(), "file-1.txt")
	assert.Equal(t, h.GetItems()[0].GetContent(), "Test 1")
	assert.Equal(t, h.GetItems()[0].GetMimeType(), "text/plain")

	h = inboxroad.NewMessageAttachmentCollectionFromSliceMap(inboxroad.SliceStringMap{
		{"_key": "file-1.txt", "_value": "Test 1", "mimeType": "text/plain"},
		{"_key": "file-2.txt", "_value": "Test 2", "mimeType": "text/plain"},
	})

	assert.Equal(t, 0, len(h.GetItems()))
}
