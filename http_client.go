package inboxroad

import "github.com/go-resty/resty/v2"

type HTTPClientInterface interface {
	NewRequest() HTTPRequestInterface
}

type HTTPClient struct {
	httpClient *resty.Client
}

func (h *HTTPClient) NewRequest() HTTPRequestInterface {
	return &HTTPRequest{ //nolint:exhaustivestruct
		httpClient: h.httpClient,
	}
}

func NewHTTPClient(options HTTPClientOptionsInterface) *HTTPClient {
	httpClient := resty.New().
		SetTimeout(options.GetTimeout()).
		SetHostURL(options.GetHostURL()).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", "inboxroad-api/inboxroad-go")

	if options.GetAPIKey() != "" {
		httpClient.SetAuthScheme(options.GetAuthScheme()).SetAuthToken(options.GetAPIKey())
	}

	if options.GetTLSConfig() != nil {
		httpClient.SetTLSClientConfig(options.GetTLSConfig())
	}

	return &HTTPClient{
		httpClient: httpClient,
	}
}
