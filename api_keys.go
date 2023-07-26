package btcpayserver

import (
	"fmt"
	"net/http"

	"github.com/negasus/btcpayserver-go/models"
)

type APIKeys struct {
	requestFunc requestFunc
}

// Revoke https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_DeleteApiKey
func (c *APIKeys) Revoke(apiKey string) error {
	if apiKey == "" {
		return fmt.Errorf("apiKey is required")
	}

	return c.requestFunc(http.MethodDelete, "/api/v1/api-keys/"+apiKey, nil, nil, nil)
}

// RevokeForUser https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_DeleteUserApiKey
func (c *APIKeys) RevokeForUser(idOrEmail, apiKey string) error {
	if idOrEmail == "" {
		return fmt.Errorf("idOrEmail is required")
	}
	if apiKey == "" {
		return fmt.Errorf("apiKey is required")
	}

	return c.requestFunc(http.MethodDelete, "/api/v1/users/"+idOrEmail+"/api-keys/"+apiKey, nil, nil, nil)
}

// GetCurrent https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_GetCurrentApiKey
func (c *APIKeys) GetCurrent() (*models.APIKey, error) {
	res := models.APIKey{}

	err := c.requestFunc(http.MethodGet, "/api/v1/api-keys/current", nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// RevokeCurrent https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_DeleteCurrentApiKey
func (c *APIKeys) RevokeCurrent() error {
	return c.requestFunc(http.MethodDelete, "/api/v1/api-keys/current", nil, nil, nil)
}

type createApiKeyRequest struct {
	Label       string              `json:"label"`
	Permissions []models.Permission `json:"permissions"`
}

// Create https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_CreateApiKey
func (c *APIKeys) Create(label string, permissions []models.Permission) (*models.APIKey, error) {
	if label == "" {
		return nil, fmt.Errorf("label is required")
	}

	res := models.APIKey{}

	req := createApiKeyRequest{
		Label:       label,
		Permissions: permissions,
	}

	err := c.requestFunc(http.MethodPost, "/api/v1/api-keys", nil, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateForUser https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_CreateUserApiKey
func (c *APIKeys) CreateForUser(idOrEmail, label string, permissions []models.Permission) (*models.APIKey, error) {
	if idOrEmail == "" {
		return nil, fmt.Errorf("idOrEmail is required")
	}
	if label == "" {
		return nil, fmt.Errorf("label is required")
	}

	res := models.APIKey{}

	req := createApiKeyRequest{
		Label:       label,
		Permissions: permissions,
	}

	err := c.requestFunc(http.MethodPost, "/api/v1/users/"+idOrEmail+"/api-keys", nil, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
