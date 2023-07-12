# Go library for BTCPayServer API

https://btcpayserver.org/

This is a go library for interacting with the BTCPayServer api.

API https://docs.btcpayserver.org/API/Greenfield/v1/

## Install

```bash
go get github.com/negasus/btcpayserver-go
```

## Usage

Create a new client with your server address and access token.

```go
client := btcpayserver.New("server address", "token")
```

Use the client to interact with the api.

```go
client.APIKeys.GetCurrent()
client.Invoices.Get(storeID, invoiceID)
client.Webhooks.Delete(storeID, webhookID)

etc...
```

## Serve Webhooks

```go
package main

import (
    "context"
    "fmt"
    "net"
	"os"
	"os/signal"
    
    "github.com/negasus/btcpayserver-go"
    "github.com/negasus/btcpayserver-go/models"
)

func main() {
	client := btcpayserver.New("server address", "token")

	ln, errLn := net.Listen("tcp", ":2000")
	if errLn != nil {
		panic(errLn)
	}
	defer ln.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	webhookSecret := "123"

	err := client.Webhooks.Serve(ctx, ln, webhookSecret, handler)
	if err != nil {
		panic(err)
	}
}

func handler(event *models.WebhookEvent) {
	fmt.Printf("Got webhook event: %+v\n", event)
}
```

## Status

At the moment, a small part of the methods has been implemented. I will be adding more methods as I need them. If you need a method that is not implemented, feel free to open an issue or a pull request.
