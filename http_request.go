package inboxroad

import (
	"fmt" //nolint:gofumpt,gci
	"github.com/go-resty/resty/v2"
)

type HTTPRequestInterface interface {
	SetBody(body string) HTTPRequestInterface
	GetBody() string
	Get(url string) (HTTPResponseInterface, error)
	Post(url string) (HTTPResponseInterface, error)
}

type HTTPRequest struct {
	httpClient *resty.Client
	body       string
}

func (h *HTTPRequest) SetBody(body string) HTTPRequestInterface {
	h.body = body

	return h
}

func (h *HTTPRequest) GetBody() string {
	return h.body
}

func (h *HTTPRequest) Get(url string) (HTTPResponseInterface, error) {
	response, err := h.httpClient.R().Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to make the http request: %w", err)
	}

	return NewHTTPResponse().
		SetIsSuccess(response.IsSuccess()).
		SetBody(response.String()).
		SetStatusCode(response.StatusCode()), nil
}

func (h *HTTPRequest) Post(url string) (HTTPResponseInterface, error) {
	response, err := h.httpClient.R().SetBody(h.body).Post(url)
	if err != nil {
		return nil, fmt.Errorf("unable to make the http request: %w", err)
	}

	return NewHTTPResponse().
		SetIsSuccess(response.IsSuccess()).
		SetBody(response.String()).
		SetStatusCode(response.StatusCode()), nil
}
