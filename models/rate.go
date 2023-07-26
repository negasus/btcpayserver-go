package models

type Rate struct {
	CurrencyPair string   `json:"currencyPair"`
	Errors       []string `json:"errors"`
	Rate         string   `json:"rate"`
}
