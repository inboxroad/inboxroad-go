package inboxroad_test

import (
	"fmt" //nolint:goimports,gofumpt,gci
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:goimports,gofumpt,gci
)

func TestMessageHeader_SetKey(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeader("", "")
	assert.Implements(t, (*inboxroad.MessageHeaderInterface)(nil), h.SetKey("X-Test"))
	assert.IsType(t, &inboxroad.MessageHeader{}, h.SetKey("X-Test"))
}

func TestMessageHeader_GetKey(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeader("", "")
	assert.Equal(t, "X-Test", h.SetKey("X-Test").GetKey())
}

func TestMessageHeader_SetValue(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeader("", "")
	assert.Implements(t, (*inboxroad.MessageHeaderInterface)(nil), h.SetValue("Test"))
	assert.IsType(t, &inboxroad.MessageHeader{}, h.SetValue("Test"))
}

func TestMessageHeader_GetValue(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeader("", "")
	assert.Equal(t, "Test", h.SetValue("Test").GetValue())
}

func TestMessageHeader_ToArray(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeader("X-Test", "Test")
	data := h.ToArray()

	type strData struct {
		Key, Value string
	}

	stringValues := []strData{
		{"key", "X-Test"},
		{"value", "Test"},
	}

	for _, strVal := range stringValues {
		if val, ok := data[strVal.Key]; ok {
			assert.Equal(t, strVal.Value, val)
		} else {
			assert.Error(t, fmt.Errorf("%w: %s", errMissingMapKey, strVal.Key))
		}
	}
}

func TestMessageHeader_ToInboxroadArray(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeader("X-Test", "Test")
	data := h.ToInboxroadArray()

	if val, ok := data["X-Test"]; ok {
		assert.Equal(t, "Test", val)
	} else {
		assert.Error(t, fmt.Errorf("%w: %s", errMissingMapKey, "X-Test"))
	}
}

func TestNewMessageHeader(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessageHeader("", "")
	assert.Implements(t, (*inboxroad.MessageHeaderInterface)(nil), h)
	assert.IsType(t, &inboxroad.MessageHeader{}, h)
}
