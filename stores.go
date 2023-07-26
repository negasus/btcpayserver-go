package btcpayserver

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/negasus/btcpayserver-go/models"
)

type Stores struct {
	requestFunc requestFunc
}

// Rates https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Stores_GetStoreRates
func (c *Stores) Rates(storeID string, currencyPairs []string) ([]models.Rate, error) {
	if storeID == "" {
		return nil, fmt.Errorf("storeID is required")
	}

	query := url.Values{}
	for _, pair := range currencyPairs {
		query.Add("currencyPair", pair)
	}

	var res []models.Rate

	err := c.requestFunc(http.MethodGet, "/api/v1/stores/"+storeID+"/rates", query, nil, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
