```go 
// Create the options for the http client 
httpClientOptions := inboxroad.NewHttpClientOptions().SetApiKey("api-key-here")

// Create the http client with the above options
httpClient := inboxroad.NewHttpClient(httpClientOptions)

// Optional, needed if you will be using ir.Messages().Send()
ir := inboxroad.NewInboxroad(httpClient)

// Build the message like: 
messageFromBuilder := NewMessage().
    setFromName("").
    setToName("...").
    setToEmail("...")

// OR like:     
messageFromMap := map[string]interface{}{
    "fromName": "...",
    "toName": "...",
    "toEmail": "...",
}

response, err := ir.Messages().Send(messageFromBuilder); 
// OR: 
// response, err := ir.Messages().Send(messageFromMap);
// OR: 
// response, err := NewMessagesApi(httpClient).Send(messageFromBuilder)
// OR: 
// response, err := NewMessagesApi(httpClient).Send(messageFromMap)

if err != nil {
   fmt.Println(err)
}

if response.GetIsSuccess() {
    fmt.Println(response.GetMessageId())
}

```