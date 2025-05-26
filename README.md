# Inboxroad Library for Go

## Getting Started
You will need an [inboxroad account](https://www.inboxroad.com) to get started.  
Once you get an account, you will need to [get your api key]([https://www.inboxroad.com/](https://app.inboxroad.com/settings/connect-api))
to use it in the API calls.

## Installation

`go get github.com/inboxroad/inboxroad-go`

## Usage

#### Send email, method 1:
```go
package main 

// Import the package
import (
    "fmt"
    "github.com/inboxroad/inboxroad-go"
    "os"
)

// Create the message object
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

// Create the http client:
httpOptions := inboxroad.NewHTTPClientOptions().SetAPIKey(os.Getenv("INBOXROAD_API_KEY"))
httpClient := inboxroad.NewHTTPClient(httpOptions)

// Create the endpoint connection
messages := inboxroad.NewMessagesAPI(httpClient)

// Send the message
response, err := messages.Send(message)

// Error?
if err != nil {
    panic(err)
}

// Get the status and the message id
fmt.Println(response.GetIsSuccess(), response.GetMessageID())
```

#### Send email, method 2:
```go
package main 

// Import the package
import (
    "fmt"
    "github.com/inboxroad/inboxroad-go"
    "os"
)

// Create the message 
message := inboxroad.NewMessageFromMap(inboxroad.StringAnyMap{
    "fromEmail":    os.Getenv("INBOXROAD_SEND_EMAIL_FROM_EMAIL"),
    "fromName":     "Inboxroad Go Client - Test Suite",
    "toEmail":      os.Getenv("INBOXROAD_SEND_EMAIL_TO_EMAIL"),
    "toName":       "Inboxroad Tester",
    "replyToEmail": os.Getenv("INBOXROAD_SEND_EMAIL_FROM_EMAIL"),
    "subject":      "Inboxroad test",
    "text":         "This is a test sent via the Inboxroad Go Client",
    "html":         "<b>This is a test sent via the Inboxroad Go Client</b>",
    "headers": inboxroad.SliceStringMap{
        {
            "key":   "X-Test1",
            "value": "Test 1",
        },
        {
            "key":   "X-Test2",
            "value": "Test 2",
        },
    },
    "attachments": inboxroad.SliceStringMap{
        {
            "name":     "test-1.txt",
            "content":  "Test 1",
            "mimeType": "text/plain",
        },
        {
            "name":     "test-2.txt",
            "content":  "Test 2",
            "mimeType": "text/plain",
        },
    },
})

// Create the http client:
httpOptions := inboxroad.NewHTTPClientOptions().SetAPIKey(os.Getenv("INBOXROAD_API_KEY"))
httpClient := inboxroad.NewHTTPClient(httpOptions)

// Create the object instance
ir := inboxroad.NewInboxroad(httpClient)

// Send the message
response, err := ir.NewMessagesAPI().Send(message)

// Error?
if err != nil {
    panic(err)
}

// Get the status and the message id
fmt.Println(response.GetIsSuccess(), response.GetMessageID())
```

## License
MIT

## Test
Following environment variable must be set in order to test the actual sending process:  
`INBOXROAD_API_KEY` - The API key for accessing the API  
`INBOXROAD_SEND_EMAIL_ENABLED` - Whether the tests should send emails (1 | 0)  
`INBOXROAD_SEND_EMAIL_FROM_EMAIL` - The email address from where the emails come from  
`INBOXROAD_SEND_EMAIL_TO_EMAIL` - The email address where emails will go  
Without these, the tests will run but no email will ever be sent.

Run the tests with:
```bash
$ go test ./...
``` 

## Bug Reports
Report [here](https://github.com/inboxroad/inboxroad-go/issues).
