package inboxroad_test

import ( //nolint:nolintlint,goimports,gofumpt,gci
	"errors" //nolint:gofumpt
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:nolintlint,goimports,gofumpt,gci
)

var errMissingMapKey = errors.New("missing required map key")

func TestNewInboxroad(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewInboxroad(inboxroad.NewHTTPClient(inboxroad.NewHTTPClientOptions()))
	assert.Implements(t, (*inboxroad.InboxroadInterface)(nil), h)
	assert.IsType(t, &inboxroad.Inboxroad{}, h)
}

func TestInboxroad_NewMessagesAPI(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewInboxroad(inboxroad.NewHTTPClient(inboxroad.NewHTTPClientOptions())).NewMessagesAPI()
	assert.Implements(t, (*inboxroad.MessagesAPIInterface)(nil), h)
	assert.IsType(t, &inboxroad.MessagesAPI{}, h)
}
