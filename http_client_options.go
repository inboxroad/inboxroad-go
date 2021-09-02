package inboxroad

import (
	"crypto/tls"
	"time"
)

type HTTPClientOptionsInterface interface {
	SetHostURL(url string) HTTPClientOptionsInterface
	GetHostURL() string
	SetAuthScheme(scheme string) HTTPClientOptionsInterface
	GetAuthScheme() string
	SetAPIKey(apiKey string) HTTPClientOptionsInterface
	GetAPIKey() string
	SetTimeout(timeout time.Duration) HTTPClientOptionsInterface
	GetTimeout() time.Duration
	SetTLSConfig(config *tls.Config) HTTPClientOptionsInterface
	GetTLSConfig() *tls.Config
}

type HTTPClientOptions struct {
	hostURL    string
	authScheme string
	apiKey     string
	timeout    time.Duration
	tlsConfig  *tls.Config
}

func (h *HTTPClientOptions) SetHostURL(url string) HTTPClientOptionsInterface {
	h.hostURL = url

	return h
}

func (h HTTPClientOptions) GetHostURL() string {
	return h.hostURL
}

func (h *HTTPClientOptions) SetAuthScheme(scheme string) HTTPClientOptionsInterface {
	h.authScheme = scheme

	return h
}

func (h HTTPClientOptions) GetAuthScheme() string {
	return h.authScheme
}

func (h *HTTPClientOptions) SetAPIKey(apiKey string) HTTPClientOptionsInterface {
	h.apiKey = apiKey

	return h
}

func (h HTTPClientOptions) GetAPIKey() string {
	return h.apiKey
}

func (h *HTTPClientOptions) SetTimeout(timeout time.Duration) HTTPClientOptionsInterface {
	h.timeout = timeout

	return h
}

func (h HTTPClientOptions) GetTimeout() time.Duration {
	return h.timeout
}

func (h *HTTPClientOptions) SetTLSConfig(config *tls.Config) HTTPClientOptionsInterface {
	h.tlsConfig = config

	return h
}

func (h HTTPClientOptions) GetTLSConfig() *tls.Config {
	return h.tlsConfig
}

func NewHTTPClientOptions() HTTPClientOptionsInterface {
	return &HTTPClientOptions{ //nolint:exhaustivestruct
		hostURL:    "https://webapi.inboxroad.com/api/v1",
		authScheme: "Basic",
		apiKey:     "",
		timeout:    time.Second * 30, //nolint:gomnd
	}
}
