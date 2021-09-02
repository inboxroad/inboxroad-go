package inboxroad

type HTTPResponseInterface interface {
	SetBody(body string) HTTPResponseInterface
	GetBody() string
	SetStatusCode(statusCode int) HTTPResponseInterface
	GetStatusCode() int
	SetIsSuccess(isSuccess bool) HTTPResponseInterface
	GetIsSuccess() bool
}

type HTTPResponse struct {
	body       string
	statusCode int
	isSuccess  bool
}

func (h *HTTPResponse) SetBody(body string) HTTPResponseInterface {
	h.body = body

	return h
}

func (h HTTPResponse) GetBody() string {
	return h.body
}

func (h *HTTPResponse) SetStatusCode(statusCode int) HTTPResponseInterface {
	h.statusCode = statusCode

	return h
}

func (h HTTPResponse) GetStatusCode() int {
	return h.statusCode
}

func (h *HTTPResponse) SetIsSuccess(isSuccess bool) HTTPResponseInterface {
	h.isSuccess = isSuccess

	return h
}

func (h HTTPResponse) GetIsSuccess() bool {
	return h.isSuccess
}

func NewHTTPResponse() HTTPResponseInterface {
	return &HTTPResponse{} //nolint:exhaustivestruct
}
