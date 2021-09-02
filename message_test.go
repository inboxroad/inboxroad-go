package inboxroad_test

import ( //nolint:nolintlint,goimports,gofumpt,gci
	"github.com/inboxroad/inboxroad-go"
	"github.com/stretchr/testify/assert"
	"testing" //nolint:nolintlint,goimports,gofumpt,gci
)

func TestMessage_SetMessageID(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetMessageID("123-abc"))
	assert.IsType(t, &inboxroad.Message{}, m.SetMessageID("123-abc"))
}

func TestMessage_GetMessageID(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Equal(t, "12345", m.SetMessageID("12345").GetMessageID())
}

func TestMessage_SetFromEmail(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetFromEmail("test@example.com"))
	assert.IsType(t, &inboxroad.Message{}, m.SetFromEmail("test@example.com"))
}

func TestMessage_GetFromEmail(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Equal(t, "test@example.com", m.SetFromEmail("test@example.com").GetFromEmail())
}

func TestMessage_SetFromName(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetFromName("test"))
	assert.IsType(t, &inboxroad.Message{}, m.SetFromName("test"))
}

func TestMessage_GetFromName(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Equal(t, "test", m.SetFromName("test").GetFromName())
}

func TestMessage_SetToEmail(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetToEmail("test@example.com"))
	assert.IsType(t, &inboxroad.Message{}, m.SetToEmail("test@example.com"))
}

func TestMessage_GetToEmail(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Equal(t, "test@example.com", m.SetToEmail("test@example.com").GetToEmail())
}

func TestMessage_SetToName(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetToName("test"))
	assert.IsType(t, &inboxroad.Message{}, m.SetToName("test"))
}

func TestMessage_GetToName(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Equal(t, "test", m.SetToName("test").GetToName())
}

func TestMessage_SetReplyToEmail(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetReplyToEmail("test@example.com"))
	assert.IsType(t, &inboxroad.Message{}, m.SetReplyToEmail("test@example.com"))
}

func TestMessage_GetReplyToEmail(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Equal(t, "test@example.com", m.SetReplyToEmail("test@example.com").GetReplyToEmail())
}

func TestMessage_SetSubject(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetSubject("subject"))
	assert.IsType(t, &inboxroad.Message{}, m.SetSubject("subject"))
}

func TestMessage_GetSubject(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Equal(t, "subject", m.SetSubject("subject").GetSubject())
}

func TestMessage_SetText(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetText("text"))
	assert.IsType(t, &inboxroad.Message{}, m.SetText("text"))
}

func TestMessage_GetText(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Equal(t, "text", m.SetText("text").GetText())
}

func TestMessage_SetHTML(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetHTML("html"))
	assert.IsType(t, &inboxroad.Message{}, m.SetHTML("html"))
}

func TestMessage_GetHTML(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	assert.Equal(t, "html", m.SetHTML("html").GetHTML())
}

func TestMessage_SetHeaders(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	h := inboxroad.NewMessageHeaderCollection()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetHeaders(h))
	assert.IsType(t, &inboxroad.Message{}, m.SetHeaders(h))
}

func TestMessage_GetHeaders(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	h := inboxroad.NewMessageHeaderCollection()
	assert.Equal(t, h, m.SetHeaders(h).GetHeaders())
}

func TestMessage_SetAttachments(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	h := inboxroad.NewMessageAttachmentCollection()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), m.SetAttachments(h))
	assert.IsType(t, &inboxroad.Message{}, m.SetAttachments(h))
}

func TestMessage_GetAttachments(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessage()
	h := inboxroad.NewMessageAttachmentCollection()
	assert.Equal(t, h, m.SetAttachments(h).GetAttachments())
}

func TestMessage_ToMap(t *testing.T) { //nolint:paralleltest
	message := inboxroad.NewMessage().
		SetFromEmail("from@example.com").
		SetFromName("Sender").
		SetToEmail("to@example.com").
		SetToName("Recipient").
		SetReplyToEmail("reply@example.com").
		SetSubject("Subject").
		SetText("Text").
		SetHTML("<b>Html</b>").
		SetHeaders(
			inboxroad.NewMessageHeaderCollection().
				Add(inboxroad.NewMessageHeader("X-Test1", "Test 1")).
				Add(inboxroad.NewMessageHeader("X-Test2", "Test 2")),
		).
		SetAttachments(
			inboxroad.NewMessageAttachmentCollection().
				Add(inboxroad.NewMessageAttachment("file-1.txt", "Test 1", "text/plain")).
				Add(inboxroad.NewMessageAttachment("file-2.txt", "Test 2", "text/plain")),
		)

	data := message.ToMap()
	result := inboxroad.StringAnyMap{
		"fromEmail":    "from@example.com",
		"fromName":     "Sender",
		"toEmail":      "to@example.com",
		"toName":       "Recipient",
		"replyToEmail": "reply@example.com",
		"subject":      "Subject",
		"text":         "Text",
		"html":         "<b>Html</b>",
		"headers": inboxroad.SliceStringMap{
			{"key": "X-Test1", "value": "Test 1"},
			{"key": "X-Test2", "value": "Test 2"},
		},
		"attachments": inboxroad.SliceStringMap{
			{"name": "file-1.txt", "content": "Test 1", "mimeType": "text/plain"},
			{"name": "file-2.txt", "content": "Test 2", "mimeType": "text/plain"},
		},
	}

	assert.Equal(t, data, result)
}

func TestMessage_ToInboxroadMap(t *testing.T) { //nolint:paralleltest
	message := inboxroad.NewMessage().
		SetFromEmail("from@example.com").
		SetFromName("Sender").
		SetToEmail("to@example.com").
		SetToName("Recipient").
		SetReplyToEmail("reply@example.com").
		SetSubject("Subject").
		SetText("Text").
		SetHTML("<b>Html</b>").
		SetHeaders(
			inboxroad.NewMessageHeaderCollection().
				Add(inboxroad.NewMessageHeader("X-Test1", "Test 1")).
				Add(inboxroad.NewMessageHeader("X-Test2", "Test 2")),
		).
		SetAttachments(
			inboxroad.NewMessageAttachmentCollection().
				Add(inboxroad.NewMessageAttachment("file-1.txt", "Test 1", "text/plain")).
				Add(inboxroad.NewMessageAttachment("file-2.txt", "Test 2", "text/plain")),
		)

	data := message.ToInboxroadMap()

	result := inboxroad.StringAnyMap{
		"from_email":     "from@example.com",
		"from_name":      "Sender",
		"to_email":       "to@example.com",
		"to_name":        "Recipient",
		"reply_to_email": "reply@example.com",
		"subject":        "Subject",
		"text":           "Text",
		"html":           "<b>Html</b>",
		"headers": inboxroad.StringMap{
			"X-Test1": "Test 1",
			"X-Test2": "Test 2",
		},
		"attachments": inboxroad.SliceStringMap{
			inboxroad.StringMap{"filename": "file-1.txt", "file_data": "VGVzdCAx", "mime_type": "text/plain"},
			inboxroad.StringMap{"filename": "file-2.txt", "file_data": "VGVzdCAy", "mime_type": "text/plain"},
		},
	}

	assert.Equal(t, data, result)
}

func TestNewMessage(t *testing.T) { //nolint:paralleltest
	h := inboxroad.NewMessage()
	assert.Implements(t, (*inboxroad.MessageInterface)(nil), h)
	assert.IsType(t, &inboxroad.Message{}, h)
}
