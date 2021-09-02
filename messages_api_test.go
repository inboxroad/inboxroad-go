package inboxroad_test

import ( //nolint:nolintlint,gofumpt,gci
	"github.com/inboxroad/inboxroad-go" //nolint:goimports
	"github.com/stretchr/testify/assert"
	"os"      //nolint:nolintlint,gofumpt,gci
	"testing" //nolint:goimports
)

func TestMessagesAPI_Send(t *testing.T) { //nolint:paralleltest
	type checker func() bool

	envKeys := []checker{
		func() bool { return os.Getenv("INBOXROAD_SEND_EMAIL_ENABLED") == "1" },
		func() bool { return len(os.Getenv("INBOXROAD_API_KEY")) == 40 },
		func() bool { return len(os.Getenv("INBOXROAD_SEND_EMAIL_FROM_EMAIL")) > 6 },
		func() bool { return len(os.Getenv("INBOXROAD_SEND_EMAIL_TO_EMAIL")) > 6 },
	}

	send := true
	for _, check := range envKeys { //nolint:wsl
		if !check() {
			send = false

			break
		}
	}

	if !send {
		t.Skip()
	}

	message := inboxroad.NewMessage().
		SetFromEmail(os.Getenv("INBOXROAD_SEND_EMAIL_FROM_EMAIL")).
		SetFromName("Inboxroad Go Client - Test Suite").
		SetToEmail(os.Getenv("INBOXROAD_SEND_EMAIL_TO_EMAIL")).
		SetToName("Inboxroad Tester").
		SetReplyToEmail(os.Getenv("INBOXROAD_SEND_EMAIL_FROM_EMAIL")).
		SetSubject("Inboxroad test").
		SetText("This is a test sent via the Inboxroad Go Client").
		SetHTML("<b>This is a test sent via the Inboxroad Go Client</b>").
		SetHeaders(
			inboxroad.NewMessageHeaderCollection().
				Add(inboxroad.NewMessageHeader("X-Test3", "Test 3")).
				Add(inboxroad.NewMessageHeader("X-Test4", "Test 4")),
		).
		SetAttachments(
			inboxroad.NewMessageAttachmentCollection().
				Add(inboxroad.NewMessageAttachment("test-3.txt", "Test 3", "text/plain")).
				Add(inboxroad.NewMessageAttachment("test-4.txt", "Test 4", "text/plain")),
		)

	messages := inboxroad.NewMessagesAPI(
		inboxroad.NewHTTPClient(
			inboxroad.NewHTTPClientOptions().
				SetAPIKey(os.Getenv("INBOXROAD_API_KEY")),
		),
	)
	response, err := messages.Send(message)

	assert.NoError(t, err)
	assert.True(t, response.GetIsSuccess())
	assert.NotEmpty(t, response.GetMessageID())
}

func TestNewMessagesAPI(t *testing.T) { //nolint:paralleltest
	m := inboxroad.NewMessagesAPI(inboxroad.NewHTTPClient(inboxroad.NewHTTPClientOptions()))
	assert.Implements(t, (*inboxroad.MessagesAPIInterface)(nil), m)
	assert.IsType(t, &inboxroad.MessagesAPI{}, m)
}
