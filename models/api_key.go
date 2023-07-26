package models

type APIKey struct {
	APIKey      string       `json:"apiKey"`
	Label       string       `json:"label"`
	Permissions []Permission `json:"permissions"`
}
