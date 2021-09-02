package inboxroad_test

import ( //nolint:nolintlint,gci
	"crypto/tls" //nolint:goimports,gofumpt
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:nolintlint,goimports,gofumpt,gci
	"time"
)

func TestHTTPClientOptions_SetHostURL(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPClientOptions()
	assert.Implements(t, (*inboxroad.HTTPClientOptionsInterface)(nil), h.SetHostURL("https://www.example.com"))
	assert.IsType(t, &inboxroad.HTTPClientOptions{}, h.SetHostURL("https://www.example.com"))
}

func TestHTTPClientOptions_GetHostURL(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPClientOptions()
	assert.Equal(t, "https://www.example.com", h.SetHostURL("https://www.example.com").GetHostURL())
}

func TestHTTPClientOptions_SetAuthScheme(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPClientOptions()
	assert.Implements(t, (*inboxroad.HTTPClientOptionsInterface)(nil), h.SetAuthScheme("Basic"))
	assert.IsType(t, &inboxroad.HTTPClientOptions{}, h.SetAuthScheme("Basic"))
}

func TestHTTPClientOptions_GetAuthScheme(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPClientOptions()
	assert.Equal(t, "Basic", h.SetAuthScheme("Basic").GetAuthScheme())
}

func TestHTTPClientOptions_SetAPIKey(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPClientOptions()
	assert.Implements(t, (*inboxroad.HTTPClientOptionsInterface)(nil), h.SetAPIKey(""))
	assert.IsType(t, &inboxroad.HTTPClientOptions{}, h.SetAPIKey(""))
}

func TestHTTPClientOptions_GetAPIKey(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPClientOptions()
	assert.Equal(t, "Test", h.SetAPIKey("Test").GetAPIKey())
}

func TestHTTPClientOptions_SetTimeout(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPClientOptions()
	assert.Implements(t, (*inboxroad.HTTPClientOptionsInterface)(nil), h.SetTimeout(time.Second))
	assert.IsType(t, &inboxroad.HTTPClientOptions{}, h.SetTimeout(time.Second))
}

func TestHTTPClientOptions_GetTimeout(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPClientOptions()
	assert.Equal(t, time.Second, h.SetTimeout(time.Second).GetTimeout())
}

func TestHTTPClientOptions_SetTLSConfig(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPClientOptions()
	assert.Implements(t, (*inboxroad.HTTPClientOptionsInterface)(nil), h.SetTLSConfig(&tls.Config{}))
	assert.IsType(t, &inboxroad.HTTPClientOptions{}, h.SetTLSConfig(&tls.Config{})) //nolint:exhaustivestruct,gosec
}

func TestHTTPClientOptions_GetTLSConfig(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewHTTPClientOptions()
	assert.Equal(t, &tls.Config{}, h.SetTLSConfig(&tls.Config{}).GetTLSConfig()) //nolint:exhaustivestruct,gosec
}
