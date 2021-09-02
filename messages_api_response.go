package inboxroad

import (
	"encoding/json"
	"strings"
)

type MessagesAPISendResponseInterface interface {
	HTTPResponseInterface
	GetMessageID() string
}

type MessagesAPISendResponse struct {
	HTTPResponse
}

func (m MessagesAPISendResponse) GetMessageID() string {
	msgResponse := struct {
		MessageID string `json:"message_id"` //nolint:tagliatelle
	}{
		MessageID: "",
	}

	err := json.Unmarshal([]byte(m.GetBody()), &msgResponse)
	if err != nil {
		return ""
	}

	return strings.Trim(msgResponse.MessageID, "<>")
}

func NewMessagesAPISendResponse() MessagesAPISendResponseInterface {
	return &MessagesAPISendResponse{
		HTTPResponse{}, //nolint:exhaustivestruct
	}
}
