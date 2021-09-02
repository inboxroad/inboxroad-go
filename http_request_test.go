package inboxroad_test

import (
	"encoding/json"
	"fmt" //nolint:goimports,gofumpt,gci
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:goimports,gofumpt,gci
)

func TestHTTPRequest_SetBody(t *testing.T) { //nolint:paralleltest
	h := &inboxroad.HTTPRequest{}
	assert.Implements(t, (*inboxroad.HTTPRequestInterface)(nil), h.SetBody(`{}`))
	assert.IsType(t, &inboxroad.HTTPRequest{}, h.SetBody(`{}`))
}

func TestHTTPRequest_GetBody(t *testing.T) { //nolint:paralleltest
	h := &inboxroad.HTTPRequest{}
	h.SetBody(`{}`)
	assert.Equal(t, `{}`, h.GetBody())
}

func TestHTTPRequest_Get(t *testing.T) {
	t.Parallel()

	url := "https://httpbin.org"
	options := inboxroad.NewHTTPClientOptions().
		SetHostURL(url).
		SetAuthScheme("").
		SetAPIKey("")

	h := inboxroad.NewHTTPClient(options)
	response, err := h.NewRequest().Get("/get")

	assert.NoError(t, err)

	d := map[string]string{
		"url": "",
	}

	err = json.Unmarshal([]byte(response.GetBody()), &d)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, fmt.Sprintf("%s/get", url), d["url"])
}

func TestHTTPRequest_Post(t *testing.T) {
	t.Parallel()

	url := "https://httpbin.org"
	options := inboxroad.NewHTTPClientOptions().
		SetHostURL(url).
		SetAuthScheme("").
		SetAPIKey("")

	h := inboxroad.NewHTTPClient(options)
	response, err := h.NewRequest().SetBody(`{}`).Post("/post")

	assert.NoError(t, err)

	d := map[string]string{
		"url": "",
	}

	err = json.Unmarshal([]byte(response.GetBody()), &d)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, fmt.Sprintf("%s/post", url), d["url"])
}
