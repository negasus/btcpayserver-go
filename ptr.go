package btcpayserver

import "github.com/negasus/btcpayserver-go/models"

func FloatPtr(v float64) *models.Float {
	f := models.Float(v)
	return &f
}

func StringPtr(v string) *string {
	return &v
}
