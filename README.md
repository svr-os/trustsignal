# TrustSignal Go Library

A lightweight Go client for the TrustSignal SMS API.

## Installation

```bash
go get github.com/svr-os/trustsignal/v1
```

## Usage

### Initialization

```go
import "github.com/svr-os/trustsignal/v1/api"

client := api.NewClient(api.WithAPIKey("YOUR_API_KEY"), api.WithBaseURL("https://api.trustsignal.io/v1"))
```

### Sending SMS

#### Bulk Send

```go
data := api.BulkSMSData{
    SenderID: "TSGNAL",
    Receivers: []api.BulkReceiver{
        {
            To:      919999999999,
            Message: "Welcome to TrustSignal. This is a test SMS.",
        },
        // ... other receivers
    },
    Route: "transactional",
    Pr1:   "customerparam1",
}
response, err := client.BulkSendSMS(data)
if err != nil {
    // handle error
}
// TODO: Use the response
```

#### Single Send

```go
data := api.SingleSMSData{
    SenderID: "TSGNAL",
    To:       []string{"919999999999"},
    Route:    "transactional",
    Message:  "Welcome to TrustSignal. This is a test SMS.",
    Pr1:      "customerparam1",
}
response, err := client.SendSMS(data)
if err != nil {
    // handle error
}
// TODO: Use the response
```

### Headers

#### Get Header

```go
response, err := client.GetHeaders()
if err != nil {
    // handle error
}
// TODO: Use the response
```

#### Create Header

```go
data := api.CreateHeaderData{
    Code:        "ABCDEF",
    DLTEntityID: "1234567890987654321",
    Company:     "company",
    Brand:       "brand",
}
response, err := client.CreateHeader(data)
if err != nil {
    // handle error
}
// TODO: Use the response
```

### Templates

#### Get Templates

```go
response, err := client.GetTemplates()
if err != nil {
    // handle error
}
// TODO: Use the response
```

#### Create Template

```go
data := api.CreateTemplateData{
    Name:        "Test Template",
    Content:     "Template Content",
    Headers:     []string{"ABCDEF", "QWERTY"},
    DLTEntityID: "1234567890987654321",
}
response, err := client.CreateTemplate(data)
if err != nil {
    // handle error
}
// TODO: Use the response
```

### Sub-Accounts

```go
data := api.CreateSubAccountData{
    Email:    "example@email.com",
    Phone:    9999999999,
    Name:     "John Doe",
    Company:  "Company",
    Password: "password123",
}
response, err := client.CreateSubAccount(data)
if err != nil {
    // handle error
}
// TODO: Use the response
```

### Webhook

To process the incoming webhook, decode the payload into the `WebhookData` struct:

```go
var data api.WebhookData
err := json.Unmarshal(yourPayload, &data)
if err != nil {
    // handle error
}
// TODO: Process the webhook data as needed
```

## Logging

The library provides a default logger. However, you can plug in your custom logger by implementing the `Logger` interface:

```go
type CustomLogger struct{}

func (l *CustomLogger) Info(msg string, keysAndValues ...interface{}) {
    // Your custom info logging logic here
}

func (l *CustomLogger) Error(msg string, keysAndValues ...interface{}) {
    // Your custom error logging logic here
}

client := api.NewClient(api.WithLogger(&CustomLogger{}))
```

---