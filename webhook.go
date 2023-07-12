package btcpayserver

import (
	"fmt"
	"net/http"

	"github.com/negasus/btcpayserver-go/models"
)

type Webhook struct {
	requestFunc requestFunc
}

// List https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Webhooks_GetWebhooks
func (c *Webhook) List(storeID string) ([]*models.Webhook, error) {
	if storeID == "" {
		return nil, fmt.Errorf("storeID is required")
	}

	var res []*models.Webhook

	err := c.requestFunc(http.MethodGet, "/api/v1/stores/"+storeID+"/webhooks", nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Webhooks_CreateWebhook
func (c *Webhook) Create(storeID string, webhook models.NewWebhook) (*models.Webhook, error) {
	if storeID == "" {
		return nil, fmt.Errorf("storeID is required")
	}

	res := models.Webhook{}

	err := c.requestFunc(http.MethodPost, "/api/v1/stores/"+storeID+"/webhooks", nil, webhook, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Webhooks_GetWebhook
func (c *Webhook) Get(storeID, webhookID string) (*models.Webhook, error) {
	if storeID == "" {
		return nil, fmt.Errorf("storeID is required")
	}
	if webhookID == "" {
		return nil, fmt.Errorf("webhookID is required")
	}

	res := models.Webhook{}

	err := c.requestFunc(http.MethodGet, "/api/v1/stores/"+storeID+"/webhooks/"+webhookID, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Update https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Webhooks_UpdateWebhook
func (c *Webhook) Update(storeID, webhookID string, webhook models.NewWebhook) (*models.Webhook, error) {
	if storeID == "" {
		return nil, fmt.Errorf("storeID is required")
	}
	if webhookID == "" {
		return nil, fmt.Errorf("webhookID is required")
	}

	res := models.Webhook{}

	err := c.requestFunc(http.MethodPut, "/api/v1/stores/"+storeID+"/webhooks/"+webhookID, nil, webhook, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Delete https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Webhooks_DeleteWebhook
func (c *Webhook) Delete(storeID, webhookID string) error {
	if storeID == "" {
		return fmt.Errorf("storeID is required")
	}
	if webhookID == "" {
		return fmt.Errorf("webhookID is required")
	}

	return c.requestFunc(http.MethodDelete, "/api/v1/stores/"+storeID+"/webhooks/"+webhookID, nil, nil, nil)
}

// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Webhooks_GetWebhookDeliveries
// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Webhooks_GetWebhookDelivery
// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Webhooks_GetWebhookDeliveryRequests
// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Webhooks_RedeliverWebhookDelivery
