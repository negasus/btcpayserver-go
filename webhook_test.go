package btcpayserver

import (
	"testing"
)

func Test_checkSig(t *testing.T) {
	body := []byte(`{
  "deliveryId": "MexLY6gqtTsH2Tj67Hs8tE",
  "webhookId": "2d6jyK97Mb2Kh95Mwfao2K",
  "originalDeliveryId": "MexLY6gqtTsH2Tj67Hs8tE",
  "isRedelivery": false,
  "type": "InvoiceCreated",
  "timestamp": 1689171308,
  "storeId": "9DMyti5DsvJRkmQnkeqdQ15mUPZDdmBtUb6j45MAGJo2",
  "invoiceId": "53e1v3oyB18WmQDrH29mUj",
  "metadata": {}
}`)
	sig := "sha256=36edeacd744b12d9e8619361b9b7674d7df1cd7e566749d3994d8e0a7341c038"
	secret := "123"

	ok := checkSig(sig, secret, body)
	if !ok {
		t.Error("checkSig failed")
	}
}
