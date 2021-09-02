package inboxroad

type (
	StringMap      = map[string]string //nolint:nolintlint,gofumpt
	SliceStringMap = []StringMap
	StringAnyMap   = map[string]interface{}
)

type InboxroadInterface interface {
	NewMessagesAPI() MessagesAPIInterface
}

type Inboxroad struct {
	httpClient HTTPClientInterface
}

func NewInboxroad(httpClient HTTPClientInterface) *Inboxroad {
	return &Inboxroad{
		httpClient: httpClient,
	}
}

func (ir Inboxroad) NewMessagesAPI() MessagesAPIInterface {
	return NewMessagesAPI(ir.httpClient)
}
