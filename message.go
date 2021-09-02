package inboxroad

type MessageInterface interface {
	SetMessageID(messageID string) MessageInterface
	GetMessageID() string
	SetFromEmail(email string) MessageInterface
	GetFromEmail() string
	SetFromName(name string) MessageInterface
	GetFromName() string
	SetToEmail(email string) MessageInterface
	GetToEmail() string
	SetToName(name string) MessageInterface
	GetToName() string
	SetReplyToEmail(email string) MessageInterface
	GetReplyToEmail() string
	SetSubject(subject string) MessageInterface
	GetSubject() string
	SetText(text string) MessageInterface
	GetText() string
	SetHTML(subject string) MessageInterface
	GetHTML() string
	SetHeaders(headers MessageHeaderCollectionInterface) MessageInterface
	GetHeaders() MessageHeaderCollectionInterface
	SetAttachments(attachments MessageAttachmentCollectionInterface) MessageInterface
	GetAttachments() MessageAttachmentCollectionInterface
	ToArray() map[string]interface{}
	ToInboxroadArray() map[string]interface{}
}

type Message struct {
	messageID    string
	fromEmail    string
	fromName     string
	toEmail      string
	toName       string
	replyToEmail string
	subject      string
	text         string
	html         string
	headers      MessageHeaderCollectionInterface
	attachments  MessageAttachmentCollectionInterface
}

func (m *Message) SetMessageID(messageID string) MessageInterface {
	m.messageID = messageID

	return m
}

func (m Message) GetMessageID() string {
	return m.messageID
}

func (m *Message) SetFromEmail(email string) MessageInterface {
	m.fromEmail = email

	return m
}

func (m Message) GetFromEmail() string {
	return m.fromEmail
}

func (m *Message) SetFromName(name string) MessageInterface {
	m.fromName = name

	return m
}

func (m Message) GetFromName() string {
	return m.fromName
}

func (m *Message) SetToEmail(email string) MessageInterface {
	m.toEmail = email

	return m
}

func (m Message) GetToEmail() string {
	return m.toEmail
}

func (m *Message) SetToName(name string) MessageInterface {
	m.toName = name

	return m
}

func (m Message) GetToName() string {
	return m.toName
}

func (m *Message) SetReplyToEmail(email string) MessageInterface {
	m.replyToEmail = email

	return m
}

func (m Message) GetReplyToEmail() string {
	return m.replyToEmail
}

func (m *Message) SetSubject(subject string) MessageInterface {
	m.subject = subject

	return m
}

func (m Message) GetSubject() string {
	return m.subject
}

func (m *Message) SetText(text string) MessageInterface {
	m.text = text

	return m
}

func (m Message) GetText() string {
	return m.text
}

func (m *Message) SetHTML(html string) MessageInterface {
	m.html = html

	return m
}

func (m Message) GetHTML() string {
	return m.html
}

func (m *Message) SetHeaders(headers MessageHeaderCollectionInterface) MessageInterface {
	m.headers = headers

	return m
}

func (m Message) GetHeaders() MessageHeaderCollectionInterface {
	return m.headers
}

func (m *Message) SetAttachments(attachments MessageAttachmentCollectionInterface) MessageInterface {
	m.attachments = attachments

	return m
}

func (m Message) GetAttachments() MessageAttachmentCollectionInterface {
	return m.attachments
}

func (m Message) ToArray() map[string]interface{} {
	return map[string]interface{}{
		"fromEmail":    m.GetFromEmail(),
		"fromName":     m.GetFromName(),
		"toEmail":      m.GetToEmail(),
		"toName":       m.GetToName(),
		"replyToEmail": m.GetReplyToEmail(),
		"subject":      m.GetSubject(),
		"text":         m.GetText(),
		"html":         m.GetHTML(),
		"headers":      m.GetHeaders().ToArray(),
		"attachments":  m.GetAttachments().ToArray(),
	}
}

func (m Message) ToInboxroadArray() map[string]interface{} {
	return map[string]interface{}{
		"from_email":     m.GetFromEmail(),
		"from_name":      m.GetFromName(),
		"to_email":       m.GetToEmail(),
		"to_name":        m.GetToName(),
		"reply_to_email": m.GetReplyToEmail(),
		"subject":        m.GetSubject(),
		"text":           m.GetText(),
		"html":           m.GetHTML(),
		"headers":        m.GetHeaders().ToInboxroadArray(),
		"attachments":    m.GetAttachments().ToInboxroadArray(),
	}
}

func NewMessage() MessageInterface {
	return &Message{ //nolint:exhaustivestruct
		headers:     NewMessageHeaderCollection(),
		attachments: NewMessageAttachmentCollection(),
	}
}

func NewMessageFromArray(params map[string]interface{}) MessageInterface {
	message := NewMessage()

	stringData := map[string]func(value string){
		"fromEmail":    func(value string) { message.SetFromEmail(value) },
		"fromName":     func(value string) { message.SetFromName(value) },
		"toEmail":      func(value string) { message.SetToEmail(value) },
		"toName":       func(value string) { message.SetToName(value) },
		"replyToEmail": func(value string) { message.SetReplyToEmail(value) },
		"subject":      func(value string) { message.SetSubject(value) },
		"text":         func(value string) { message.SetText(value) },
		"html":         func(value string) { message.SetHTML(value) },
	}

	for key, callback := range stringData {
		if value, ok := params[key].(string); ok {
			callback(value)
		}
	}

	if rawHeaders, ok := params["headers"].([]map[string]string); ok {
		headers := NewMessageHeaderCollectionFromArray(rawHeaders)
		for _, header := range headers.GetItems() {
			message.GetHeaders().Add(header)
		}
	}

	if rawAttachments, ok := params["attachments"].([]map[string]string); ok {
		attachments := NewMessageAttachmentCollectionFromArray(rawAttachments)
		for _, attachment := range attachments.GetItems() {
			message.GetAttachments().Add(attachment)
		}
	}

	return message
}
