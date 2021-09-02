package inboxroad_test

import ( //nolint:nolintlint,goimports,gofumpt,gci
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:nolintlint,goimports,gofumpt,gci
)

func TestMessageHeaderCollection_Add(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeaderCollection().
		Add(inboxroad.NewMessageHeader("X-Test1", "Test 1")).
		Add(inboxroad.NewMessageHeader("X-Test2", "Test 2"))

	assert.Implements(t, (*inboxroad.MessageHeaderCollectionInterface)(nil), h)
	assert.IsType(t, &inboxroad.MessageHeaderCollection{}, h)
}

func TestMessageHeaderCollection_GetItems(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeaderCollection().
		Add(inboxroad.NewMessageHeader("X-Test1", "Test 1")).
		Add(inboxroad.NewMessageHeader("X-Test2", "Test 2"))

	items := h.GetItems()
	assert.Equal(t, 2, len(items))
	assert.Equal(t, items[0].GetKey(), "X-Test1")
	assert.Equal(t, items[0].GetValue(), "Test 1")
}

func TestMessageHeaderCollection_ToArray(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeaderCollection().
		Add(inboxroad.NewMessageHeader("X-Test1", "Test 1")).
		Add(inboxroad.NewMessageHeader("X-Test2", "Test 2"))

	data := h.ToArray()

	result := []map[string]string{
		{"key": "X-Test1", "value": "Test 1"},
		{"key": "X-Test2", "value": "Test 2"},
	}

	assert.Equal(t, result, data)
}

func TestMessageHeaderCollection_ToInboxroadArray(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeaderCollection().
		Add(inboxroad.NewMessageHeader("X-Test1", "Test 1")).
		Add(inboxroad.NewMessageHeader("X-Test2", "Test 2"))

	data := h.ToInboxroadArray()

	result := map[string]string{
		"X-Test1": "Test 1",
		"X-Test2": "Test 2",
	}

	assert.Equal(t, result, data)
}

func TestNewMessageHeaderCollection(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeaderCollection()
	assert.Implements(t, (*inboxroad.MessageHeaderCollectionInterface)(nil), h)
	assert.IsType(t, &inboxroad.MessageHeaderCollection{}, h)
}

func TestNewMessageHeaderCollectionFromArray(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeaderCollectionFromArray([]map[string]string{
		{"key": "X-Test1", "value": "Test 1"},
		{"key": "X-Test2", "value": "Test 2"},
	})

	assert.Equal(t, 2, len(h.GetItems()))
	assert.Equal(t, h.GetItems()[0].GetKey(), "X-Test1")
	assert.Equal(t, h.GetItems()[0].GetValue(), "Test 1")

	h = inboxroad.NewMessageHeaderCollectionFromArray([]map[string]string{
		{"_key": "X-Test1", "_value": "Test 1"},
		{"_key": "X-Test2", "_value": "Test 2"},
	})

	assert.Equal(t, 0, len(h.GetItems()))
}