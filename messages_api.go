package inboxroad

import (
	"encoding/json"
	"fmt"
)

type MessagesAPIInterface interface {
	Send(message MessageInterface) (MessagesAPISendResponseInterface, error)
}

type MessagesAPI struct {
	httpClient HTTPClientInterface
}

func (m *MessagesAPI) Send(message MessageInterface) (MessagesAPISendResponseInterface, error) {
	jsonBody, err := json.Marshal(message.ToInboxroadMap())
	if err != nil {
		return nil, fmt.Errorf("invalid message object: %w", err)
	}

	response, err := m.httpClient.
		NewRequest().
		SetBody(string(jsonBody)).
		Post("/messages/")

	if err != nil {
		return nil, fmt.Errorf("error while doing the http request: %w", err)
	}

	resp := NewMessagesAPISendResponse()
	resp.
		SetIsSuccess(response.GetIsSuccess()).
		SetBody(response.GetBody()).
		SetStatusCode(response.GetStatusCode())

	return resp, nil
}

func NewMessagesAPI(httpClient HTTPClientInterface) *MessagesAPI {
	return &MessagesAPI{httpClient: httpClient}
}
