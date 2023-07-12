package models

type AuthorizedEvents struct {
	Everything     bool               `json:"everything"`
	SpecificEvents []WebhookEventType `json:"specificEvents"`
}

type Webhook struct {
	Enabled             bool             `json:"enabled"`
	AutomaticRedelivery bool             `json:"automaticRedelivery"`
	URL                 string           `json:"url"`
	AuthorizedEvents    AuthorizedEvents `json:"authorizedEvents"`
	ID                  string           `json:"id"`
	Secret              *string          `json:"secret"`
}

type NewWebhook struct {
	Enabled             bool             `json:"enabled"`
	AutomaticRedelivery bool             `json:"automaticRedelivery"`
	URL                 string           `json:"url"`
	AuthorizedEvents    AuthorizedEvents `json:"authorizedEvents"`
	Secret              *string          `json:"secret"`
}
