package inboxroad_test

import ( //nolint:nolintlint,goimports,gofumpt,gci
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:nolintlint,goimports,gofumpt,gci
)

func TestHTTPResponse_SetBody(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPResponse()
	assert.Implements(t, (*inboxroad.HTTPResponseInterface)(nil), h.SetBody(""))
	assert.IsType(t, &inboxroad.HTTPResponse{}, h.SetBody(""))
}

func TestHTTPResponse_GetBody(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPResponse()
	h.SetBody("Test")
	assert.Equal(t, "Test", h.GetBody())
}

func TestHTTPResponse_SetStatusCode(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPResponse()
	assert.Implements(t, (*inboxroad.HTTPResponseInterface)(nil), h.SetStatusCode(0))
	assert.IsType(t, &inboxroad.HTTPResponse{}, h.SetStatusCode(0))
}

func TestHTTPResponse_GetStatusCode(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPResponse()
	h.SetStatusCode(200)
	assert.Equal(t, 200, h.GetStatusCode())
}

func TestHTTPResponse_SetIsSuccess(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPResponse()
	assert.Implements(t, (*inboxroad.HTTPResponseInterface)(nil), h.SetIsSuccess(true))
	assert.IsType(t, &inboxroad.HTTPResponse{}, h.SetIsSuccess(true))
}

func TestHTTPResponse_GetIsSuccess(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPResponse()
	h.SetIsSuccess(true)
	assert.True(t, h.GetIsSuccess())
}
