package inboxroad_test

import ( //nolint:goimports,gofumpt,nolintlint
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:goimports,gofumpt,nolintlint
)

func TestNewHTTPClient(t *testing.T) { //nolint:paralleltest
	assert.Implements(t, (*inboxroad.HTTPClientInterface)(nil), inboxroad.NewHTTPClient(inboxroad.NewHTTPClientOptions()))
}

func TestHTTPClient_NewRequest(t *testing.T) { //nolint:paralleltest
	assert.Implements(t,
		(*inboxroad.HTTPRequestInterface)(nil),
		inboxroad.NewHTTPClient(inboxroad.NewHTTPClientOptions()).NewRequest(),
	)
}
