package models

type WebhookEventType string

const (
	WebhookEventTypeInvoiceCreated         WebhookEventType = "InvoiceCreated"
	WebhookEventTypeInvoiceReceivedPayment WebhookEventType = "InvoiceReceivedPayment"
	WebhookEventTypeInvoiceProcessing      WebhookEventType = "InvoiceProcessing"
	WebhookEventTypeInvoiceExpired         WebhookEventType = "InvoiceExpired"
	WebhookEventTypeInvoiceSettled         WebhookEventType = "InvoiceSettled"
	WebhookEventTypeInvoiceInvalid         WebhookEventType = "InvoiceInvalid"
	WebhookEventTypeInvoicePaymentSettled  WebhookEventType = "InvoicePaymentSettled"
)

type WebhookEvent struct {
	DeliveryID         string           `json:"deliveryId"`
	WebhookID          string           `json:"webhookId"`
	OriginalDeliveryID string           `json:"originalDeliveryId"`
	IsRedelivery       bool             `json:"isRedelivery"`
	Type               WebhookEventType `json:"type"`
	Timestamp          Time             `json:"timestamp"`
	StoreID            string           `json:"storeId"`
	InvoiceID          string           `json:"invoiceId"`
	Metadata           Metadata         `json:"metadata"`
	PartiallyPaid      bool             `json:"partiallyPaid"`   // InvoiceExpired
	AfterExpiration    bool             `json:"afterExpiration"` // InvoiceReceivedPayment, InvoicePaymentSettled
	PaymentMethod      string           `json:"paymentMethod"`   // InvoiceReceivedPayment, InvoicePaymentSettled
	Payment            *Payment         `json:"payment"`         // InvoiceReceivedPayment, InvoicePaymentSettled
	OverPaid           bool             `json:"overPaid"`        // InvoiceProcessing, InvoiceSettled
	ManuallyMarked     bool             `json:"manuallyMarked"`  // InvoiceInvalid, InvoiceSettled
}
