package models

type ReceiptOptions struct {
	Enabled      *bool `json:"enabled"`
	ShowQR       *bool `json:"showQR"`
	ShowPayments *bool `json:"showPayments"`
}

type SpeedPolicy string

const (
	SpeedPolicyHighSpeed      SpeedPolicy = "HighSpeed"
	SpeedPolicyLowMediumSpeed SpeedPolicy = "LowMediumSpeed"
	SpeedPolicyLowSpeed       SpeedPolicy = "LowSpeed"
	SpeedPolicyMediumSpeed    SpeedPolicy = "MediumSpeed"
)

type CheckoutType string

const (
	CheckoutTypeV1 CheckoutType = "V1"
	CheckoutTypeV2 CheckoutType = "V2"
)

type CheckoutOptions struct {
	SpeedPolicy           SpeedPolicy   `json:"speedPolicy"`
	PaymentMethods        []string      `json:"paymentMethods"`
	DefaultPaymentMethod  *string       `json:"defaultPaymentMethod"`
	LazyPaymentMethods    *bool         `json:"lazyPaymentMethods"`
	ExpirationMinutes     *int          `json:"expirationMinutes"`
	MonitoringMinutes     *int          `json:"monitoringMinutes"`
	PaymentTolerance      *float64      `json:"paymentTolerance"`
	RedirectURL           *string       `json:"redirectURL"`
	RedirectAutomatically *bool         `json:"redirectAutomatically"`
	RequiresRefundEmail   *bool         `json:"requiresRefundEmail"`
	CheckoutType          *CheckoutType `json:"checkoutType"`
	DefaultLanguage       *string       `json:"defaultLanguage"`
}

type InvoiceType string

const (
	InvoiceTypeStandard InvoiceType = "Standard"
	InvoiceTypeTopUp    InvoiceType = "TopUp"
)

type InvoiceStatus string

const (
	InvoiceStatusExpired    InvoiceStatus = "expired"
	InvoiceStatusInvalid    InvoiceStatus = "invalid"
	InvoiceStatusNew        InvoiceStatus = "new"
	InvoiceStatusProcessing InvoiceStatus = "processing"
	InvoiceStatusSettled    InvoiceStatus = "settled"
)

type InvoiceAdditionalStatus string

const (
	InvoiceAdditionalStatusInvalid     InvoiceAdditionalStatus = "Invalid"
	InvoiceAdditionalStatusMarked      InvoiceAdditionalStatus = "Marked"
	InvoiceAdditionalStatusNone        InvoiceAdditionalStatus = "None"
	InvoiceAdditionalStatusPaidLate    InvoiceAdditionalStatus = "PaidLate"
	InvoiceAdditionalStatusPaidOver    InvoiceAdditionalStatus = "PaidOver"
	InvoiceAdditionalStatusPaidPartial InvoiceAdditionalStatus = "PaidPartial"
)

type Invoice struct {
	Metadata                          Metadata        `json:"metadata"`
	Checkout                          CheckoutOptions `json:"checkout"`
	Receipt                           ReceiptOptions  `json:"receipt"`
	ID                                string          `json:"id"`
	StoreID                           string          `json:"storeId"`
	Amount                            Float           `json:"amount"`
	Currency                          string          `json:"currency"`
	Type                              InvoiceType     `json:"type"`
	CheckoutLink                      string          `json:"checkoutLink"`
	CreatedTime                       Time            `json:"createdTime"`
	ExpirationTime                    Time            `json:"expirationTime"`
	MonitoringExpiration              Time            `json:"monitoringExpiration"`
	Status                            InvoiceStatus   `json:"status"`
	AdditionalStatus                  string          `json:"additionalStatus"`
	AvailableStatusesForManualMarking []InvoiceStatus `json:"availableStatusesForManualMarking"`
	Archived                          bool            `json:"archived"`
}

type NewInvoice struct {
	Metadata              *Metadata        `json:"metadata,omitempty"`
	Checkout              *CheckoutOptions `json:"checkout,omitempty"`
	Receipt               *ReceiptOptions  `json:"receipt,omitempty"`
	Amount                *Float           `json:"amount,omitempty"`
	Currency              *string          `json:"currency,omitempty"`
	AdditionalSearchTerms []string         `json:"additionalSearchTerms,omitempty"`
}
