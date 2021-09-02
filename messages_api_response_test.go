package inboxroad_test

import ( //nolint:nolintlint,goimports,gofumpt,gci
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:goimports,gofumpt
)

func TestMessagesAPISendResponse_GetMessageID(t *testing.T) { //nolint:paralleltest
	r := inboxroad.NewMessagesAPISendResponse()
	r.SetBody(`{"message_id":"12345-56789"}`)
	assert.Equal(t, "12345-56789", r.GetMessageID())
}

func TestNewMessagesAPISendResponse(t *testing.T) { //nolint:paralleltest
	r := inboxroad.NewMessagesAPISendResponse()
	assert.Implements(t, (*inboxroad.MessagesAPISendResponseInterface)(nil), r)
	assert.IsType(t, &inboxroad.MessagesAPISendResponse{}, r) //nolint:exhaustivestruct
}
