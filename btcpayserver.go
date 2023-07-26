package btcpayserver

import (
	"net/http"
	"net/url"
)

// https://docs.btcpayserver.org/API/Greenfield/v1/

type requestFunc func(method, path string, query url.Values, payload any, dest any) error

type Client struct {
	server string
	token  string

	httpClient *http.Client

	APIKeys  *APIKeys
	Invoices *Invoices
	Stores   *Stores
	Webhooks *Webhook
}

func New(server, token string) *Client {
	c := &Client{
		server: server,
		token:  token,

		httpClient: http.DefaultClient,
	}

	c.APIKeys = &APIKeys{requestFunc: c.request}
	c.Invoices = &Invoices{requestFunc: c.request}
	c.Stores = &Stores{requestFunc: c.request}
	c.Webhooks = &Webhook{requestFunc: c.request}

	return c
}
