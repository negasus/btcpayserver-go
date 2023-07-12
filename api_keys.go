package btcpayserver

import (
	"net/http"

	"github.com/negasus/btcpayserver-go/models"
)

type APIKeys struct {
	requestFunc requestFunc
}

// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_DeleteApiKey
// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_DeleteUserApiKey

// GetCurrent https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_GetCurrentApiKey
func (c *APIKeys) GetCurrent() (*models.APIKey, error) {
	res := models.APIKey{}

	err := c.requestFunc(http.MethodGet, "/api/v1/api-keys/current", nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_DeleteCurrentApiKey
// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_CreateApiKey
// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/ApiKeys_CreateUserApiKey
